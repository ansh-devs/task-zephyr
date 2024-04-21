package internal

import (
	"context"
	pb "github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (s *Orchestrator) UpdateTaskStatus(ctx context.Context, req *pb.UpdateTaskStatusRequest) (*pb.UpdateTaskStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskStatus not implemented")
}

func (s *Orchestrator) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	s.WorkerPoolMtx.Lock()
	defer s.WorkerPoolMtx.Unlock()
	// id is used to fetch the worker id.
	id := req.GetWorkerId()
	// addr contains the worker address
	addr := req.GetAddress()
	if worker, ok := s.WorkerPool[id]; ok {
		// worker is running as intended.
		worker.HealthCheckMiss = 0
	} else {
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logrus.WithFields(logrus.Fields{"workerID": id, "workerAddress": addr}).Error("failed to register the worker")
			return &pb.HealthCheckResponse{
				IsRegistered: false,
			}, err
		}
		//Adding the worker to the worker pool.
		s.WorkerPool[id] = &Worker{
			HealthCheckMiss: 0,
			Address:         addr,
			Manager:         conn,
			Recipient:       pb.NewBackgroundWorkerServiceClient(conn),
		}
		s.AcquirableWorkerIDsMtx.Lock()
		defer s.AcquirableWorkerIDsMtx.Unlock()
		currentWorkers := len(s.WorkerPool)
		s.AcquirableWorkerIDs = make([]string, 0, currentWorkers)
		for k, _ := range s.WorkerPool {
			s.AcquirableWorkerIDs = append(s.AcquirableWorkerIDs, k)
		}
		logrus.WithFields(logrus.Fields{"workerID": id, "workerAddress": addr}).Info("worker registered successfully")

	}
	//return nil, status.Errorf(codes.Unimplemented, "method health check unimplmented.")
	return &pb.HealthCheckResponse{
		IsRegistered: true,
	}, nil
}
