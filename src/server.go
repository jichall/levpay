package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func serve(host, port string) {
	router := mux.NewRouter()

	router.HandleFunc("/super/", handleSuper).Methods("POST", "DELETE")
	// param can be an uuid or a name
	router.HandleFunc("/super/{param}", handleSearch).Methods("GET")
	// factions can be: heroes|evils|nil
	router.HandleFunc("/supers/{faction}", handleSupers).Methods("GET")

	server := http.Server{
		Addr:    host + ":" + port,
		Handler: router,
	}

	erro := make(chan error)

	go func() {
		erro <- server.ListenAndServe()
	}()

	logger.Info("server initializing on %s:%s", host, port)

	for {
		select {
		case err := <-erro:
			logger.Error("Failed to initialize server. Reason %v", err)
			os.Exit(1)
		}
	}
}
