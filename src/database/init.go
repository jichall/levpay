package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rafaelcn/jeolgyu"
)

var (
	pool *pgxpool.Pool

	logger *jeolgyu.Jeolgyu
)

// Init stats the database logger and also the schema for the application
func Init(production string) {
	var err error
	tp := jeolgyu.SinkBoth

	if len(production) > 0 {
		tp = jeolgyu.SinkFile
	}

	logger, err = jeolgyu.New(jeolgyu.Settings{
		Filepath: "log",
		Filename: "database",
		SinkType: tp,
	})

	if err != nil {
		log.Fatalf("failed to initialize database logger. Reason: %v", err)
	}
}

// Connect sets up the connection with the database
func Connect(url string) {
	if len(url) == 0 {
		logger.Warning("missing DATABASE_URL, can't connect")
		os.Exit(0)
	}

	var err error
	pool, err = pgxpool.Connect(context.Background(), url)

	if err != nil {
		logger.Error("failed to connect to the database, reason %v", err)
	}

	check()
}

// Close ends the connection to the database
func Close() {
	pool.Close()
}

// check looks for connection issues
func check() {
	conn, err := pool.Acquire(context.Background())

	if err != nil {
		logger.Error("connection acquirement failed, reason %v", err)
	} else {
		ctx, cancel := context.WithDeadline(context.Background(),
			time.Now().Add(2*time.Second))
		defer cancel()

		err = conn.Conn().Ping(ctx)

		if err != nil {
			logger.Error("database pool is unhealthy, reason %v", err)
			os.Exit(1)
		} else {
			logger.Info("database connection ok")
		}
	}
}
