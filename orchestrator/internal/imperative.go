package internal

import (
	"context"
	"github.com/sirupsen/logrus"
)

func (s *Orchestrator) Start(ctx context.Context) error {
	go s.managePool()
	err := s.StartServer()
	if err != nil {
		logrus.Error("failed to start the rpc server...")
	}
	//s.DataStorePool, err = ConnectToDatabase("", ctx)
	go s.ScrapeDatabaseForJobs()
	return s.gracefulShutdown()
}
