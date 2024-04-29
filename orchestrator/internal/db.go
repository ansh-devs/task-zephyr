package internal

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"time"
)

func ConnectToDatabase(dbString string, ctx context.Context) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	dbtries := 0
	tkr := time.NewTicker(time.Second * 1)
	for range tkr.C {
		pool, err := pgxpool.New(ctx, dbString)

		if err == nil {
			tkr.Stop()
			return pool, nil
		} else {
			logrus.Error(err)
			dbtries++
		}

	}
	return pool, nil
}
