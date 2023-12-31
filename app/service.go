package app

import (
	"context"
	"fmt"
	"nashrul-be/crm/utils/logutils"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
)

var envPath string

func openElog(name string) debug.Log {
	elog, err := eventlog.Open(name)
	if err != nil {
		return nil
	}
	return elog
}

type myservice struct{}

func (m *myservice) Execute(_ []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	elog := openElog("be-portal-tro")
	defer elog.Close()
	changes <- svc.Status{State: svc.StartPending}
	srv := Init(envPath)
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
				defer cancel()
				if err := srv.Shutdown(ctx); err != nil {
					logutils.Get().Printf("Server shutdown: %v\n", err)
				}
				select {
				case <-ctx.Done():
					logutils.Get().Println("timeout of 1 second.")
				}
				logutils.Get().Println("Server exiting")
				break loop
			default:
				elog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func RunService(name, env string) {
	envPath = env
	elog, err := eventlog.Open(name)
	if err != nil {
		return
	}
	defer elog.Close()
	elog.Info(1, fmt.Sprintf("starting %s service", name))
	run := svc.Run
	err = run(name, &myservice{})
	if err != nil {
		elog.Error(1, fmt.Sprintf("%s service failed: %v", name, err))
		return
	}
	elog.Info(1, fmt.Sprintf("%s service stopped", name))
}
