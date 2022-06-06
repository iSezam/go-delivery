// Package app: OS shutdown hook
package app

import (
	"os"
	"os/signal"
	"syscall"
)

var osSignal chan os.Signal

func init() {
	osSignal = make(chan os.Signal, 1)
	signal.Notify(
		osSignal,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
}

func onShutdown(actionHook func()) {
	<-osSignal
	actionHook()
}
