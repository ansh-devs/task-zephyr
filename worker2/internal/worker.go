package internal

import (
	"context"
	"github.com/ansh-devs/task-zephyr/worker/protov3/protos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Worker struct {
	protos.UnimplementedBackgroundWorkerServiceServer
	ln             net.Listener
	grpcSrvr       *grpc.Server
	Port           string
	HealthCheckTTL int
	workerID       string
	Ctx            context.Context
	CtxCancel      context.CancelFunc
}

func NewWorker(port string) *Worker {
	var worker Worker
	worker.Port = port
	worker.workerID = uuid.NewString()
	worker.HealthCheckTTL = 1
	ln, err := ListenToAddr(port)
	if err != nil {
		logrus.Fatalf("error while net.Listen : %v", err)
	}
	worker.ln = ln
	worker.SetUp()
	return &worker
}

func ListenToAddr(port string) (net.Listener, error) {
	return net.Listen("tcp", port)
}

func (w *Worker) SetUp() {
	w.grpcSrvr = grpc.NewServer()
	protos.RegisterBackgroundWorkerServiceServer(w.grpcSrvr, w)
	reflection.Register(w.grpcSrvr)

}

func (w *Worker) Serve() {
	if err := w.grpcSrvr.Serve(w.ln); err != nil {
		logrus.Fatalf("error while starting rpc server : %v", err)
	}
}
