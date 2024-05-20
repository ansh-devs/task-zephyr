package internal

import (
	"context"

	"github.com/ansh-devs/task-zephyr/worker/protov3/protos"
	"github.com/sirupsen/logrus"
)

func (b *Worker) AssignTaskToWorker(ctx context.Context, req *protos.AssignTaskToWorkerRequest) (*protos.AssignTaskToWorkerResponse, error) {

	logrus.Infof("processing task with id : %s", req.GetJobId())
	return &protos.AssignTaskToWorkerResponse{
		IsAccepted: true,
	}, nil
}
