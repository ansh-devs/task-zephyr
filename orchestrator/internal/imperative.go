package internal

import (
	"context"
	"github.com/sirupsen/logrus"
)

func (o *Orchestrator) Start(ctx context.Context) error {
	go o.managePool()
	err := o.StartServer()
	if err != nil {
		logrus.Error("failed to start the rpc server...")
	}
	o.DataStorePool, err = ConnectToDatabase("", ctx)
	go o.ScrapeDatabaseForJobs()
	return o.gracefulShutdown()
}
