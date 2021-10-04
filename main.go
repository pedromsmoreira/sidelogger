package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pedromsmoreira/sidelogger/loggers"
)

func main() {
	logger := &loggers.NullLogger{}
	server := loggers.NewServer(logger)
	err := server.Start("localhost", "5000")

	if err != nil {

		os.Exit(1)
	}
	ss := make(chan os.Signal, 1)
	signal.Notify(ss, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-ss
	server.Shutdown()
	os.Exit(0)
}
