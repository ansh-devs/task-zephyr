package internal

import (
	"context"
	"github.com/ansh-devs/task-zephyr/worker/protov3/protos"
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
	Ctx            context.Context
	CtxCancel      context.CancelFunc
}

func NewWorker(port string) *Worker {
	var worker Worker
	worker.Port = port
	ln, err := ListenToAddr(port)
	if err != nil {
		logrus.Fatalf("error fetching ip address : %v", err)
	}
	worker.ln = ln
	worker.HealthCheckTTL = 1
	worker.SetUp()
	return &worker
}

func ListenToAddr(port string) (net.Listener, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	} else {
		return ln, nil
	}
}

func (b *Worker) SetUp() {
	b.grpcSrvr = grpc.NewServer()
	protos.RegisterBackgroundWorkerServiceServer(b.grpcSrvr, b)
	reflection.Register(b.grpcSrvr)

}

func (b *Worker) Serve() {
	if err := b.grpcSrvr.Serve(b.ln); err != nil {
		logrus.Fatalf("cannot start rpc server : %v", err)
	}
}
