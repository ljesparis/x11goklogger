package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func Atexit(fn func()) {
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
		<-sc

		if fn != nil {
			fn()
		}

		os.Exit(1)
	}()
}
