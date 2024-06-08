package internal

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	"github.com/ansh-devs/task-zephyr/worker/protov3/protos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func (b *Worker) SendHealthCheck() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcConn, err := grpc.Dial("localhost:50000", opts...)
	if err != nil {
		logrus.Errorf("error occured in grpc.Dial : %v", err)
	}
	workerID := uuid.NewString()
	client := protos.NewOrchestratorServiceClient(grpcConn)
	ticker := time.NewTicker(time.Duration(b.HealthCheckTTL))
	ipAddr := GetIPAddr().String()
	for range ticker.C {
		_, err := client.HealthCheck(context.Background(), &protos.HealthCheckRequest{
			WorkerId: workerID,
			Address:  ipAddr + b.Port,
		})
		if err != nil {
			logrus.Errorf("error occured while sending healthcheck %v", err)
			continue
		}
	}

	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			logrus.Errorf("error occured in clientConn.Close : %v", err)
		}
	}(grpcConn)
}
