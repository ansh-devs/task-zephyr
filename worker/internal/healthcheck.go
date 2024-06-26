package internal

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	"github.com/ansh-devs/task-zephyr/worker/protov3/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func (w *Worker) SendHealthCheck() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcConn, err := grpc.Dial("localhost:50000", opts...)
	if err != nil {
		logrus.Errorf("error occured in grpc.Dial : %v", err)
	}
	client := protos.NewOrchestratorServiceClient(grpcConn)
	ticker := time.NewTicker(time.Duration(w.HealthCheckTTL))
	ipAddr := GetIPAddr().String()
	for range ticker.C {
		_, err := client.HealthCheck(context.Background(), &protos.HealthCheckRequest{
			WorkerId: w.workerID,
			Address:  ipAddr + w.Port,
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
