package internal

import (
	"context"
	"github.com/ansh-devs/task-zephyr/worker/taskhandler"

	"github.com/ansh-devs/task-zephyr/worker/protov3/protos"
	"github.com/sirupsen/logrus"
)

func (w *Worker) AssignTaskToWorker(ctx context.Context, req *protos.AssignTaskToWorkerRequest) (*protos.AssignTaskToWorkerResponse, error) {
	logrus.WithFields(logrus.Fields{"job_id": req.GetJobId(), "job_type": req.GetJobType()}).Info("processing_task")
	if req.JobType == "SEND_MAIL" {
		if err := taskhandler.SendMailTask(req.GetCommand()); err != nil {
			return &protos.AssignTaskToWorkerResponse{
				IsAccepted: true,
				IsDone:     false,
				Error:      err.Error(),
			}, nil
		}
	} else {
	}
	return &protos.AssignTaskToWorkerResponse{IsAccepted: true}, nil
}
