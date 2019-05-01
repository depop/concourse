// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/db"
)

type FakeResourceConfigVersion struct {
	CheckOrderStub        func() int
	checkOrderMutex       sync.RWMutex
	checkOrderArgsForCall []struct {
	}
	checkOrderReturns struct {
		result1 int
	}
	checkOrderReturnsOnCall map[int]struct {
		result1 int
	}
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	MetadataStub        func() db.ResourceConfigMetadataFields
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct {
	}
	metadataReturns struct {
		result1 db.ResourceConfigMetadataFields
	}
	metadataReturnsOnCall map[int]struct {
		result1 db.ResourceConfigMetadataFields
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct {
	}
	reloadReturns struct {
		result1 bool
		result2 error
	}
	reloadReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ResourceConfigScopeStub        func() db.ResourceConfigScope
	resourceConfigScopeMutex       sync.RWMutex
	resourceConfigScopeArgsForCall []struct {
	}
	resourceConfigScopeReturns struct {
		result1 db.ResourceConfigScope
	}
	resourceConfigScopeReturnsOnCall map[int]struct {
		result1 db.ResourceConfigScope
	}
	VersionStub        func() db.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 db.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 db.Version
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceConfigVersion) CheckOrder() int {
	fake.checkOrderMutex.Lock()
	ret, specificReturn := fake.checkOrderReturnsOnCall[len(fake.checkOrderArgsForCall)]
	fake.checkOrderArgsForCall = append(fake.checkOrderArgsForCall, struct {
	}{})
	fake.recordInvocation("CheckOrder", []interface{}{})
	fake.checkOrderMutex.Unlock()
	if fake.CheckOrderStub != nil {
		return fake.CheckOrderStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkOrderReturns
	return fakeReturns.result1
}

func (fake *FakeResourceConfigVersion) CheckOrderCallCount() int {
	fake.checkOrderMutex.RLock()
	defer fake.checkOrderMutex.RUnlock()
	return len(fake.checkOrderArgsForCall)
}

func (fake *FakeResourceConfigVersion) CheckOrderCalls(stub func() int) {
	fake.checkOrderMutex.Lock()
	defer fake.checkOrderMutex.Unlock()
	fake.CheckOrderStub = stub
}

func (fake *FakeResourceConfigVersion) CheckOrderReturns(result1 int) {
	fake.checkOrderMutex.Lock()
	defer fake.checkOrderMutex.Unlock()
	fake.CheckOrderStub = nil
	fake.checkOrderReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfigVersion) CheckOrderReturnsOnCall(i int, result1 int) {
	fake.checkOrderMutex.Lock()
	defer fake.checkOrderMutex.Unlock()
	fake.CheckOrderStub = nil
	if fake.checkOrderReturnsOnCall == nil {
		fake.checkOrderReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.checkOrderReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfigVersion) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeResourceConfigVersion) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeResourceConfigVersion) IDCalls(stub func() int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeResourceConfigVersion) IDReturns(result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfigVersion) IDReturnsOnCall(i int, result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfigVersion) Metadata() db.ResourceConfigMetadataFields {
	fake.metadataMutex.Lock()
	ret, specificReturn := fake.metadataReturnsOnCall[len(fake.metadataArgsForCall)]
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct {
	}{})
	fake.recordInvocation("Metadata", []interface{}{})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.metadataReturns
	return fakeReturns.result1
}

func (fake *FakeResourceConfigVersion) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeResourceConfigVersion) MetadataCalls(stub func() db.ResourceConfigMetadataFields) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = stub
}

func (fake *FakeResourceConfigVersion) MetadataReturns(result1 db.ResourceConfigMetadataFields) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 db.ResourceConfigMetadataFields
	}{result1}
}

func (fake *FakeResourceConfigVersion) MetadataReturnsOnCall(i int, result1 db.ResourceConfigMetadataFields) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = nil
	if fake.metadataReturnsOnCall == nil {
		fake.metadataReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfigMetadataFields
		})
	}
	fake.metadataReturnsOnCall[i] = struct {
		result1 db.ResourceConfigMetadataFields
	}{result1}
}

func (fake *FakeResourceConfigVersion) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	ret, specificReturn := fake.reloadReturnsOnCall[len(fake.reloadArgsForCall)]
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct {
	}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.reloadReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceConfigVersion) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeResourceConfigVersion) ReloadCalls(stub func() (bool, error)) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = stub
}

func (fake *FakeResourceConfigVersion) ReloadReturns(result1 bool, result2 error) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigVersion) ReloadReturnsOnCall(i int, result1 bool, result2 error) {
	fake.reloadMutex.Lock()
	defer fake.reloadMutex.Unlock()
	fake.ReloadStub = nil
	if fake.reloadReturnsOnCall == nil {
		fake.reloadReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.reloadReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigVersion) ResourceConfigScope() db.ResourceConfigScope {
	fake.resourceConfigScopeMutex.Lock()
	ret, specificReturn := fake.resourceConfigScopeReturnsOnCall[len(fake.resourceConfigScopeArgsForCall)]
	fake.resourceConfigScopeArgsForCall = append(fake.resourceConfigScopeArgsForCall, struct {
	}{})
	fake.recordInvocation("ResourceConfigScope", []interface{}{})
	fake.resourceConfigScopeMutex.Unlock()
	if fake.ResourceConfigScopeStub != nil {
		return fake.ResourceConfigScopeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.resourceConfigScopeReturns
	return fakeReturns.result1
}

func (fake *FakeResourceConfigVersion) ResourceConfigScopeCallCount() int {
	fake.resourceConfigScopeMutex.RLock()
	defer fake.resourceConfigScopeMutex.RUnlock()
	return len(fake.resourceConfigScopeArgsForCall)
}

func (fake *FakeResourceConfigVersion) ResourceConfigScopeCalls(stub func() db.ResourceConfigScope) {
	fake.resourceConfigScopeMutex.Lock()
	defer fake.resourceConfigScopeMutex.Unlock()
	fake.ResourceConfigScopeStub = stub
}

func (fake *FakeResourceConfigVersion) ResourceConfigScopeReturns(result1 db.ResourceConfigScope) {
	fake.resourceConfigScopeMutex.Lock()
	defer fake.resourceConfigScopeMutex.Unlock()
	fake.ResourceConfigScopeStub = nil
	fake.resourceConfigScopeReturns = struct {
		result1 db.ResourceConfigScope
	}{result1}
}

func (fake *FakeResourceConfigVersion) ResourceConfigScopeReturnsOnCall(i int, result1 db.ResourceConfigScope) {
	fake.resourceConfigScopeMutex.Lock()
	defer fake.resourceConfigScopeMutex.Unlock()
	fake.ResourceConfigScopeStub = nil
	if fake.resourceConfigScopeReturnsOnCall == nil {
		fake.resourceConfigScopeReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfigScope
		})
	}
	fake.resourceConfigScopeReturnsOnCall[i] = struct {
		result1 db.ResourceConfigScope
	}{result1}
}

func (fake *FakeResourceConfigVersion) Version() db.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.versionReturns
	return fakeReturns.result1
}

func (fake *FakeResourceConfigVersion) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeResourceConfigVersion) VersionCalls(stub func() db.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeResourceConfigVersion) VersionReturns(result1 db.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 db.Version
	}{result1}
}

func (fake *FakeResourceConfigVersion) VersionReturnsOnCall(i int, result1 db.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 db.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 db.Version
	}{result1}
}

func (fake *FakeResourceConfigVersion) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkOrderMutex.RLock()
	defer fake.checkOrderMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.resourceConfigScopeMutex.RLock()
	defer fake.resourceConfigScopeMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceConfigVersion) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.ResourceConfigVersion = new(FakeResourceConfigVersion)