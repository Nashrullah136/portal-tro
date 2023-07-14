package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/sys/windows/svc"
	"log"
	"nashrul-be/crm/app"
	"nashrul-be/crm/utils/logutils"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var (
	envPath  string
	svcName  = "be-portal-tro"
	descSrvc = "Back End Portal TRO"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop, pause or continue.\n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func GetCurrentDir() (string, error) {
	program := os.Args[0]
	absPath, err := filepath.Abs(program)
	if err != nil {
		return "", err
	}
	return filepath.Dir(absPath), nil
}

func InitFlag() {
	cd, err := GetCurrentDir()
	if err != nil {
		panic("can't get current directory")
	}
	flag.StringVar(&envPath, "env", filepath.Join(cd, ".env"), "path to .env file")
}

func main() {
	InitFlag()
	flag.Parse()
	inService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("failed to determine if we are running in service: %v", err)
	}
	if inService {
		app.RunService(svcName, envPath)
		return
	}
	cmd := strings.ToLower(os.Args[len(os.Args)-1])
	switch cmd {
	case "debug":
		os.Setenv("LOG", "cli")
		srv := app.Init(envPath)
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logutils.Get().Println("Shutdown Server ...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logutils.Get().Fatal("Server Shutdown:", err)
		}
		select {
		case <-ctx.Done():
			logutils.Get().Println("timeout of 5 seconds.")
		}
		logutils.Get().Println("Server exiting")
	case "install":
		err = app.InstallService(svcName, descSrvc, envPath)
	case "remove":
		err = app.RemoveService(svcName)
	case "start":
		err = app.StartService(svcName, envPath)
	case "stop":
		err = app.ControlService(svcName, svc.Stop, svc.Stopped)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	}
	return
}
