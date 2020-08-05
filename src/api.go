package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handleSuper(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// TODO: read the request body and save the super
	case "DELETE":
		_ = mux.Vars(r)["uuid"]
		// TODO: read the sent param and delete the entity
	}
}

func handleSupers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if len(vars["faction"]) > 0 {
		switch vars["faction"] {
		case "heroes":
			// TODO: return only all heroes
		case "evils":
			// TODO: return only all evil stuff
		}
	} else {
		// TODO: return all superheroes
	}
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if len(vars["param"]) > 0 {
		// TODO: Filter either a name or an UUID
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
