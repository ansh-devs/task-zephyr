package internal

import (
	"context"
	pb "github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (s *Orchestrator) AssignTaskToWorker() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	tx, err := s.DataStorePool.Begin(ctx)
	if err != nil {
		logrus.Error("Error initiating the transaction in datastore")
		return
	}
	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && err == pgx.ErrTxClosed {
			logrus.Errorf("failed to perform rollback : %s", err.Error())
		}
	}()
	result, err := tx.Query(ctx, `SELECT id, scheduled_at, command, task_type, scheduled_for, type FROM jobs WHERE scheduled_at < (NOW() + INTERVAL  '20 seconds') AND started_at IS NULL ORDER BY scheduled_at FOR UPDATE SKIP LOCKED`)
	if err != nil {
		return
	}
	defer result.Close()
	var jobs []*pb.AssignTaskToWorkerRequest

	for result.Next() {
		var id, command, task_type string
		if err := result.Scan(&id, &command, &task_type); err != nil {
			logrus.Infof("failed to scan the row %v\n", err)
		}
		jobs = append(jobs, &pb.AssignTaskToWorkerRequest{
			JobId:   id,
			JobType: task_type,
			Command: command,
		})
	}

	for _, job := range jobs {
		if err := s.handleTask(); err != nil {
			logrus.Error("failed to assign task to the worker")
		}
		var workerId string
		for id, _ := range s.WorkerPool {
			workerId = id
			break
		}
		worker := s.WorkerPool[workerId]
		workerResponse, err := worker.Recipient.AssignTaskToWorker(s.Ctx, &pb.AssignTaskToWorkerRequest{
			JobId:   job.JobId,
			JobType: job.JobType,
			Command: job.Command,
		})
		if err != nil {
			logrus.Error("error while assigning task to the worker")
			return
		}
		if workerResponse.GetIsAccepted() {
			logrus.WithFields(logrus.Fields{"workerId": workerResponse.GetWorkerId(), "jobId": job.JobId}).Info("task has been assigned to a worker")
		}
	}

}
