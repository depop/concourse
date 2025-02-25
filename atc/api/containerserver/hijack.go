package containerserver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/api/accessor"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/gorilla/websocket"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

var upgrader = websocket.Upgrader{
	HandshakeTimeout: 5 * time.Second,
}

type InterceptTimeoutError struct {
	duration time.Duration
}

func (err InterceptTimeoutError) Error() string {
	return fmt.Sprintf("idle timeout (%s) reached", err.duration)
}

//counterfeiter:generate . InterceptTimeoutFactory
type InterceptTimeoutFactory interface {
	NewInterceptTimeout() InterceptTimeout
}

func NewInterceptTimeoutFactory(duration time.Duration) InterceptTimeoutFactory {
	return &interceptTimeoutFactory{
		duration: duration,
	}
}

type interceptTimeoutFactory struct {
	duration time.Duration
}

func (t *interceptTimeoutFactory) NewInterceptTimeout() InterceptTimeout {
	return &interceptTimeout{
		duration: t.duration,
		timer:    time.NewTimer(t.duration),
	}
}

//counterfeiter:generate . InterceptTimeout
type InterceptTimeout interface {
	Reset()
	Channel() <-chan time.Time
	Error() error
}

type interceptTimeout struct {
	duration time.Duration
	timer    *time.Timer
}

func (t *interceptTimeout) Reset() {
	if t.duration > 0 {
		t.timer.Reset(t.duration)
	}
}

func (t *interceptTimeout) Channel() <-chan time.Time {
	if t.duration > 0 {
		return t.timer.C
	}
	return make(chan time.Time)
}

func (t *interceptTimeout) Error() error {
	return InterceptTimeoutError{duration: t.duration}
}

func (s *Server) HijackContainer(team db.Team) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		handle := r.FormValue(":id")

		hLog := s.logger.Session("hijack", lager.Data{
			"handle": handle,
		})

		container, _, found, err := s.workerPool.LocateContainer(ctx, team.ID(), handle)
		if err != nil {
			hLog.Error("failed-to-find-container", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			hLog.Info("container-not-found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		isCheckContainer, err := team.IsCheckContainer(handle)
		if err != nil {
			hLog.Error("failed-to-find-container", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if isCheckContainer {
			acc := accessor.GetAccessor(r)
			if !acc.IsAdmin() {
				hLog.Error("user-not-authorized-to-hijack-check-container", err)
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}

		ok, err := team.IsContainerWithinTeam(handle, isCheckContainer)
		if err != nil {
			hLog.Error("failed-to-find-container-within-team", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !ok {
			hLog.Error("container-not-found-within-team", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		hLog.Debug("found-container")

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			hLog.Error("unable-to-upgrade-connection-for-websockets", err)
			return
		}

		defer db.Close(conn)

		var processSpec atc.HijackProcessSpec
		err = conn.ReadJSON(&processSpec)
		if err != nil {
			hLog.Error("malformed-process-spec", err)
			closeWithErr(hLog, conn, websocket.CloseUnsupportedData, "malformed process spec")
			return
		}

		hijackRequest := hijackRequest{
			Container: container,
			Process:   processSpec,
		}

		s.hijack(r.Context(), hLog, conn, hijackRequest)
	})
}

type hijackRequest struct {
	Container runtime.Container
	Process   atc.HijackProcessSpec
}

func closeWithErr(log lager.Logger, conn *websocket.Conn, code int, reason string) {
	err := conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(code, reason),
		time.Time{},
	)

	if err != nil {
		log.Error("failed-to-close-websocket-connection", err)
	}
}

func (s *Server) hijack(ctx context.Context, hLog lager.Logger, conn *websocket.Conn, request hijackRequest) {
	hLog = hLog.Session("hijack", lager.Data{
		"handle":  request.Container.DBContainer().Handle(),
		"process": request.Process,
	})

	stdinR, stdinW := io.Pipe()
	defer db.Close(stdinW)

	inputs := make(chan atc.HijackInput)
	outputs := make(chan atc.HijackOutput)
	exited := make(chan int, 1)
	errs := make(chan error, 1)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	outW := &stdoutWriter{
		outputs: outputs,
		done:    ctx.Done(),
	}

	errW := &stderrWriter{
		outputs: outputs,
		done:    ctx.Done(),
	}

	var tty *runtime.TTYSpec
	var idle InterceptTimeout

	if request.Process.TTY != nil {
		tty = &runtime.TTYSpec{
			WindowSize: runtime.WindowSize{
				Columns: request.Process.TTY.WindowSize.Columns,
				Rows:    request.Process.TTY.WindowSize.Rows,
			},
		}
	}

	process, err := request.Container.Run(ctx, runtime.ProcessSpec{
		Path: request.Process.Path,
		Args: request.Process.Args,
		Env:  request.Process.Env,
		Dir:  request.Process.Dir,

		User: request.Process.User,

		TTY: tty,
	}, runtime.ProcessIO{
		Stdin:  stdinR,
		Stdout: outW,
		Stderr: errW,
	})
	if err != nil {
		if errors.As(err, &runtime.ExecutableNotFoundError{}) {
			hLog.Info("executable-not-found")

			_ = conn.WriteJSON(atc.HijackOutput{
				ExecutableNotFound: true,
			})
		}

		_ = conn.WriteJSON(atc.HijackOutput{
			Error: err.Error(),
		})
		hLog.Error("failed-to-hijack", err)
		return
	}

	err = request.Container.DBContainer().UpdateLastHijack()
	if err != nil {
		hLog.Error("failed-to-update-container-hijack-time", err)
		return
	}

	go func() {
		for {
			select {
			case <-s.clock.After(s.interceptUpdateInterval):
				err = request.Container.DBContainer().UpdateLastHijack()
				if err != nil {
					hLog.Error("failed-to-update-container-hijack-time", err)
					return
				}

			case <-ctx.Done():
				return
			}
		}
	}()

	hLog.Info("hijacked")

	go func() {
		for {
			var input atc.HijackInput
			err := conn.ReadJSON(&input)
			if err != nil {
				break
			}

			select {
			case inputs <- input:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		result, err := process.Wait(ctx)
		if err != nil {
			errs <- err
		} else {
			exited <- result.ExitStatus
		}
	}()

	idle = s.interceptTimeoutFactory.NewInterceptTimeout()
	idleChan := idle.Channel()

	for {
		select {
		case input := <-inputs:
			idle.Reset()
			if input.Closed {
				_ = stdinW.Close()
			} else if input.TTYSpec != nil {
				err := process.SetTTY(runtime.TTYSpec{
					WindowSize: runtime.WindowSize{
						Columns: input.TTYSpec.WindowSize.Columns,
						Rows:    input.TTYSpec.WindowSize.Rows,
					},
				})
				if err != nil {
					_ = conn.WriteJSON(atc.HijackOutput{
						Error: err.Error(),
					})
				}
			} else {
				_, _ = stdinW.Write(input.Stdin)
			}

		case <-idleChan:
			errs <- idle.Error()

		case output := <-outputs:
			err := conn.WriteJSON(output)
			if err != nil {
				return
			}

		case status := <-exited:
			_ = conn.WriteJSON(atc.HijackOutput{
				ExitStatus: &status,
			})

			return

		case err := <-errs:
			_ = conn.WriteJSON(atc.HijackOutput{
				Error: err.Error(),
			})

			return
		}
	}
}

type stdoutWriter struct {
	outputs chan<- atc.HijackOutput
	done    <-chan struct{}
}

func (writer *stdoutWriter) Write(b []byte) (int, error) {
	chunk := make([]byte, len(b))
	copy(chunk, b)

	output := atc.HijackOutput{
		Stdout: chunk,
	}

	select {
	case writer.outputs <- output:
	case <-writer.done:
	}

	return len(b), nil
}

type stderrWriter struct {
	outputs chan<- atc.HijackOutput
	done    <-chan struct{}
}

func (writer *stderrWriter) Write(b []byte) (int, error) {
	chunk := make([]byte, len(b))
	copy(chunk, b)

	output := atc.HijackOutput{
		Stderr: chunk,
	}

	select {
	case writer.outputs <- output:
	case <-writer.done:
	}

	return len(b), nil
}
