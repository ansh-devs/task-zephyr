package main

import (
	"context"
	"fmt"

	"github.com/ansh-devs/task-zephyr/worker/protov3/protos"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcConn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		fmt.Println(err)
	}
	client := protos.NewOrchestratorServiceClient(grpcConn)
	worker, err := client.HealthCheck(context.Background(), &protos.HealthCheckRequest{
		WorkerId: "",
		Address:  "",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(worker.String())
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			logrus.Error(err)
		}
	}(grpcConn)
}
