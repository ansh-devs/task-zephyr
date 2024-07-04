package main

import (
	"github.com/ansh-devs/task-zephyr/worker/internal"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.WithFields(logrus.Fields{"status": "started", "service": "worker"}).Info("Task-Zephyr")
}

func main() {
	port := ":" + os.Getenv("PORT")
	srvc := internal.NewWorker(port)
	go srvc.SendHealthCheck()
	srvc.Serve()
}
