package internal

import (
	"context"
	pb "github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
)

type Orchestrator struct {
	pb.UnimplementedOrchestratorServiceServer
	server    *grpc.Server
	listener  net.Listener
	workerMtx sync.Mutex
}

func NewOrchestrator(server *grpc.Server, ln net.Listener) *Orchestrator {
	return &Orchestrator{
		server:   server,
		listener: ln,
	}
}

func (s *Orchestrator) AssignTaskToWorker(ctx context.Context, request *pb.AssignTaskToWorkerRequest) (*pb.AssignTaskToWorkerResponse, error) {
	logrus.Info("AssignTaskToWorker Invoked")
	return &pb.AssignTaskToWorkerResponse{
		Status: "",
	}, nil
}

func (s *Orchestrator) SaveResult(ctx context.Context, request *pb.SaveResultRequest) (*pb.SaveResultResponse, error) {
	logrus.Info("SaveResult Invoked")
	logrus.Info(request.String())
	return &pb.SaveResultResponse{
		Message: "result saved",
	}, nil
}

func (s *Orchestrator) Serve() {
	err := s.server.Serve(s.listener)
	if err != nil {
		logrus.Error(err)
	}
}

func (s *Orchestrator) PerformReflection() {
	reflection.Register(s.server)
}
