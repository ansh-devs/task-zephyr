package main

import (
	"github.com/ansh-devs/task-zephyr/worker/internal"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.WithFields(logrus.Fields{"status": "started", "service": "orchestrator"}).Info("Task-Zephyr")
}

func main() {
	srvc := internal.NewWorker(":8081")
	go srvc.SendHealthCheck()
	srvc.Serve()
}
