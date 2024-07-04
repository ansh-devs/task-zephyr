package main

import (
	"context"
	"github.com/ansh-devs/task-zephyr/orchestrator/internal"
	"github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.WithFields(log.Fields{"status": "started", "service": "orchestrator"}).Info("Task-Zephyr")
}

func main() {
	port := ":" + os.Getenv("PORT")
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Error(err)
	}
	srv := grpc.NewServer()
	orchestrator := internal.NewOrchestrator(srv, ln, port, context.Background())
	protos.RegisterOrchestratorServiceServer(srv, orchestrator)
	reflection.Register(orchestrator.Manager)
	err = orchestrator.Start(context.Background())
	if err != nil {
		log.Error(err)
	}
	orchestrator.Serve()
}
