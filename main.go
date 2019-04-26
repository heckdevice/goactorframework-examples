package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/heckdevice/goactorframework-corelib"
	"github.com/heckdevice/goactorframework-examples/samples"
)

var (
	killPill         = make(chan os.Signal)
	terminateProcess = make(chan bool)
)

func main() {
	signal.Notify(killPill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGTSTP)
	oncomingMessages := samples.InitSampleMessageQueue()
	core.GetDefaultActorSystem().Start(oncomingMessages)
	for {
		select {
		case <-killPill:
			fmt.Println(fmt.Sprintf("\n\n******--- Shutting down due to SIGTERM ---******"))
			core.GetDefaultActorSystem().Close(terminateProcess)
		case <-terminateProcess:
			fmt.Println(fmt.Sprintf("\n\n******--- Actor system is stopped, exiting ---******"))
			return
		}
	}

}
