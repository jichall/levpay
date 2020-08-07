package database

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func query(query string, arguments ...interface{}) (pgx.Rows, error) {
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		logger.Info("couldn't get a connection with the database, reason %v",
			err)
		return nil, err
	}

	defer conn.Release()

	results, err := conn.Query(context.Background(), query, arguments...)

	if err != nil {
		logger.Info("couldn't execute query, reason %v", err)
		return nil, err
	}

	return results, nil
}

func execute(query string, arguments ...interface{}) pgconn.CommandTag {
	var result pgconn.CommandTag

	conn, err := pool.Acquire(context.Background())

	if err != nil {
		logger.Error("couldn't get a pool with the database, reason %v", err)
	} else {
		defer conn.Release()

		result, err = conn.Exec(context.Background(), query, arguments...)

		if err != nil {
			logger.Error("couldn't execute query, reason %v", err)
		}
	}

	return result
}
