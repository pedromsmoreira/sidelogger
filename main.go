package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pedromsmoreira/sidelogger/loggers"
)

func main() {
	logger, _ := loggers.NewZapLogger()
	server := loggers.NewServer(loggers.NewLoggerService(logger))
	err := server.Start("4000")

	if err != nil {

		os.Exit(1)
	}
	ss := make(chan os.Signal, 1)
	logger.PlainInfo("Sidelogger is running. Ctrl-C to exit")
	signal.Notify(ss, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-ss
	logger.PlainInfo("Sidelogger is terminating!")
	err = server.Shutdown()
	if err != nil {
		logger.PlainError(fmt.Sprintf("Failed to shutdown API server: %v", err))
	}
	logger.PlainInfo("Sidelogger terminated!")
	logger.Close()
	os.Exit(0)
}
