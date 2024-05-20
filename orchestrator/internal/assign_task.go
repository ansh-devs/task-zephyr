package internal

import (
	"context"

	pb "github.com/ansh-devs/task-zephyr/orchestrator/protov3/protos"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

func (s *Orchestrator) AssignTaskToWorker() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	tx, err := s.DataStorePool.Begin(ctx)
	if err != nil {
		logrus.Errorf("Error initiating the transaction in datastore : %v", err.Error())
		return
	}
	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && err == pgx.ErrTxClosed {
			logrus.Errorf("failed to perform rollback : %s", err.Error())
		}
	}()
	result, err := tx.Query(ctx, `SELECT id, scheduled_for, command, task_type FROM jobs WHERE scheduled_for < (NOW() + INTERVAL  '20 seconds') AND started_at IS NULL ORDER BY scheduled_at FOR UPDATE SKIP LOCKED`)
	if err != nil {
		logrus.Error(err)
	}

	defer result.Close()
	var jobs []*pb.AssignTaskToWorkerRequest

	for result.Next() {
		var scheduledAt pgtype.Timestamptz
		var id, command, taskType string
		if err := result.Scan(&id, &scheduledAt, &command, &taskType); err != nil {
			logrus.Infof("failed to scan the row %v\n", err)
		}
		logrus.Infof("GOT SOME DATA id=%s, time=%s, command=%s, task-type=%s", id, scheduledAt.Time.String(), command, taskType)
		jobs = append(jobs, &pb.AssignTaskToWorkerRequest{
			JobId:   id,
			JobType: taskType,
			Command: command,
		})
	}

	for _, job := range jobs {
		if !s.areWorkersAvailable() {
			logrus.Error("no workers to process task")
			break
		}
		var workerId string
		for id := range s.WorkerPool {
			workerId = id
			break
		}
		worker := s.WorkerPool[workerId]
		logrus.Info("assigning task to the worker")
		workerResponse, err := worker.Recipient.AssignTaskToWorker(s.Ctx, &pb.AssignTaskToWorkerRequest{
			JobId:   job.JobId,
			JobType: job.JobType,
			Command: job.Command,
		})
		if err != nil {
			logrus.Errorf("error while assigning task to the worker : %v", err)
			return
		}
		if workerResponse.GetIsAccepted() {
			logrus.WithFields(logrus.Fields{"workerId": workerResponse.GetWorkerId(), "jobId": job.JobId}).Info("task has been assigned to a worker")
		}
	}

}
