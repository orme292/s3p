package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func SigIntListener() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		<-sig
		os.Exit(0)
	}()
}
