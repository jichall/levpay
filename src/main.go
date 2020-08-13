package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	db "github.com/jichall/levpay/src/database"
	"github.com/jichall/levpay/src/superapi"
	"github.com/rafaelcn/jeolgyu"
)

var (
	logger *jeolgyu.Jeolgyu
)

func main() {
	// get the default environment settings to initialize the application
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	url := os.Getenv("DATABASE_URL")

	// The API token
	token := os.Getenv("TOKEN")
	superapi.SetToken(token)

	// to prevent variable shadowing when initializing the logger
	var err error

	logger, err = jeolgyu.New(jeolgyu.Settings{
		SinkType: jeolgyu.SinkOutput,
	})

	if err != nil {
		log.Fatalf("failed to initialize default logger, reason %v", err)
	}

	db.Init()
	db.Connect(url)
	defer db.Close()

	go serve(host, port)

	// Trap exit signals. This is a good thing when we want to sanitize anything
	// in our application, one would just have to adjust the trapped signals
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Kill, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-quit:
			logger.Info("exiting, bye...")
			os.Exit(0)
		}
	}
}
