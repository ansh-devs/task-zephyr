package main

import (
	"context"
	"fmt"
	"github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
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
	defer grpcConn.Close()
}
