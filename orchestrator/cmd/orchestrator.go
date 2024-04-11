package main

import (
	"github.com/ansh-devs/task-zephyr/orchestrator/internal"
	"github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.WithFields(log.Fields{"status": "started", "service": "orchestrator"}).Info("Task-Zephyr")
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error(err)
	}
	srv := grpc.NewServer()
	orchestrator := internal.NewOrchestrator(srv, ln)
	protos.RegisterOrchestratorServiceServer(srv, orchestrator)
	orchestrator.PerformReflection()
	orchestrator.Serve()
}
