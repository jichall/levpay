package main

import (
	"log"
	"os"

	"github.com/jichall/levpay/src/database"
	"github.com/rafaelcn/jeolgyu"
)

var (
	logger *jeolgyu.Jeolgyu
)

func main() {
	// to prevent variable shadowing when initializing the logger
	var err error
	logger, err = jeolgyu.New(jeolgyu.Settings{
		Filepath: "log",
		// this decreases app performance and can be easily changed when in a
		// production environment.
		SinkType: jeolgyu.SinkBoth,
	})

	if err != nil {
		log.Fatalf("Failed to initialize default logger. Reason %v", err)
	}

	// get the default environment settings to initialize the application
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	// 4106895309380853
	token := os.Getenv("API_TOKEN")

	url := os.Getenv("DATABASE_URL")

	serve(host, port)
	database.Connect(url)
}
