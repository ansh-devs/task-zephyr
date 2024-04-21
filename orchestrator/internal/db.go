package internal

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
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
			dbtries++
		}

	}
	return pool, nil
}
