package main

import (
	"app-controller/pkg/server/v1"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
}

func main() {
	logger := logrus.StandardLogger()

	s, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	s.Start()

	stop := make(chan os.Signal, 1)
	// pkill -15 main
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	logger.Printf("Shutting down server")
	
	logrus.Exit(0)
}
