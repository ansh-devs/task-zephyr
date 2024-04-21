package grpcapi

import (
	"context"
	pb "github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	pb.UnimplementedOrchestratorServiceServer
}

func (s *GrpcServer) UpdateTaskStatus(context.Context, *pb.UpdateTaskStatusRequest) (*pb.UpdateTaskStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskStatus not implemented")
}

func (s *GrpcServer) HealthCheck(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method healthcheck unimplmented.")
}
