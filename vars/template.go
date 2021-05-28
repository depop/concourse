package vars

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type Template struct {
	bytes []byte
}

type EvaluateOpts struct {
	ExpectAllKeys bool
}

func NewTemplate(bytes []byte) Template {
	return Template{bytes: bytes}
}

func (t Template) ExtraVarNames() []string {
	return interpolator{}.extractVarNames(string(t.bytes))
}

func (t Template) Evaluate(vars Variables, opts EvaluateOpts) ([]byte, error) {
	var obj interface{}

	err := yaml.Unmarshal(t.bytes, &obj)
	if err != nil {
		return []byte{}, err
	}

	obj, err = t.interpolateRoot(obj, newVarsTracker(vars, opts.ExpectAllKeys))
	if err != nil {
		return []byte{}, err
	}

	bytes, err := yaml.Marshal(obj)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

func (t Template) interpolateRoot(obj interface{}, tracker varsTracker) (interface{}, error) {
	var err error
	obj, err = interpolator{}.Interpolate(obj, tracker)
	if err != nil {
		return nil, err
	}

	return obj, tracker.Error()
}

func ExtractVars(in interface{}) ([]Reference, error) {
	byteParams, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	tpl := NewTemplate(byteParams)
	varNames := tpl.ExtraVarNames()

	var varRefs []Reference
	for _, varName := range varNames {
		varRef, err := ParseReference(varName)
		if err != nil {
			return nil, err
		}

		varRefs = append(varRefs, varRef)
	}

	return varRefs, nil
}

type interpolator struct{}

var (
	interpolationRegex         = regexp.MustCompile(`\(\((([-/\.\w\pL]+\:)?[-/\.:@"\w\pL]+)\)\)`)
	interpolationAnchoredRegex = regexp.MustCompile("\\A" + interpolationRegex.String() + "\\z")
)

func (i interpolator) Interpolate(node interface{}, tracker varsTracker) (interface{}, error) {
	switch typedNode := node.(type) {
	case map[interface{}]interface{}:
		for k, v := range typedNode {
			evaluatedValue, err := i.Interpolate(v, tracker)
			if err != nil {
				return nil, err
			}

			evaluatedKey, err := i.Interpolate(k, tracker)
			if err != nil {
				return nil, err
			}

			delete(typedNode, k) // delete in case key has changed
			typedNode[evaluatedKey] = evaluatedValue
		}

	case []interface{}:
		for idx, x := range typedNode {
			var err error
			typedNode[idx], err = i.Interpolate(x, tracker)
			if err != nil {
				return nil, err
			}
		}

	case string:
		for _, name := range i.extractVarNames(typedNode) {
			foundVal, found, err := tracker.Get(name)
			if err != nil {
				return nil, err
			}

			if found {
				// ensure that value type is preserved when replacing the entire field
				if interpolationAnchoredRegex.MatchString(typedNode) {
					return foundVal, nil
				}

				switch foundVal.(type) {
				case string, int, int16, int32, int64, uint, uint16, uint32, uint64, json.Number:
					foundValStr := fmt.Sprintf("%v", foundVal)
					typedNode = strings.Replace(typedNode, fmt.Sprintf("((%s))", name), foundValStr, -1)
				default:
					return nil, InvalidInterpolationError{
						Name:  name,
						Value: foundVal,
					}
				}
			}
		}

		return typedNode, nil
	}

	return node, nil
}

func (i interpolator) extractVarNames(value string) []string {
	var names []string

	for _, match := range interpolationRegex.FindAllSubmatch([]byte(value), -1) {
		names = append(names, string(match[1]))
	}

	return names
}

type varsTracker struct {
	vars Variables

	expectAllFound bool

	missing    map[string]struct{}
	visited    map[string]struct{}
	visitedAll map[string]struct{} // track all var names that were accessed
}

func newVarsTracker(vars Variables, expectAllFound bool) varsTracker {
	return varsTracker{
		vars:           vars,
		expectAllFound: expectAllFound,
		missing:        map[string]struct{}{},
		visited:        map[string]struct{}{},
		visitedAll:     map[string]struct{}{},
	}
}

// Get value of a var. Name can be the following formats: 1) 'foo', where foo
// is var name; 2) 'foo:bar', where foo is var source name, and bar is var name;
// 3) '.:foo', where . means a local var, foo is var name.
func (t varsTracker) Get(varName string) (interface{}, bool, error) {
	varRef, err := ParseReference(varName)
	if err != nil {
		return nil, false, err
	}

	t.visitedAll[identifier(varRef)] = struct{}{}

	val, found, err := t.vars.Get(varRef)
	if !found || err != nil {
		t.missing[varRef.String()] = struct{}{}
		return val, found, err
	}

	return val, true, err
}

func (t varsTracker) Error() error {
	return t.MissingError()
}

func (t varsTracker) MissingError() error {
	if !t.expectAllFound || len(t.missing) == 0 {
		return nil
	}

	return UndefinedVarsError{Vars: names(t.missing)}
}

func names(mapWithNames map[string]struct{}) []string {
	var names []string
	for name, _ := range mapWithNames {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

func identifier(varRef Reference) string {
	id := varRef.Path

	if varRef.Source != "" {
		id = fmt.Sprintf("%s:%s", varRef.Source, id)
	}

	return id
}

