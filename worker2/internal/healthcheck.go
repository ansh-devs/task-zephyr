package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/ansh-devs/task-zephyr/worker/protov3/protos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (b *Worker) SendHealthCheck() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcConn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		fmt.Println(err)
	}
	workerID := uuid.NewString()
	client := protos.NewOrchestratorServiceClient(grpcConn)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		_, err := client.HealthCheck(context.Background(), &protos.HealthCheckRequest{
			WorkerId: workerID,
			Address:  GetIPAddr().String() + b.Port,
		})
		if err != nil {
			logrus.Errorf("error occured while sending healthcheck %v", err)
			continue
		}
		logrus.Info("sended heartbeat to orchestrator")

	}

	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(grpcConn)
}
