package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pedromsmoreira/sidelogger/loggers"
)

func main() {
	lf := &loggers.LoggerFactory{}
	cfg, err := loggers.ReadConfigFile("./config.yaml")
	if err != nil {
		fmt.Printf("Error occurred reading config file. Error: %v", err)
		os.Exit(1)
	}

	logger, err := lf.CreateLogger(cfg.Logger.Name)
	if err != nil {
		fmt.Printf("Error creating logger. Error: %v", err)
		os.Exit(1)
	}

	logger.PlainInfo(fmt.Sprintf("Logger %v was loaded with sucess", logger.GetName()))

	server := loggers.NewServer(loggers.NewLoggerService(logger))
	err = server.Start("4000")

	if err != nil {
		fmt.Printf("Error starting the server. Error: %v", err)
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
