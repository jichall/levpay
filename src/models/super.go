package models

type (
	// Super is an entity represint a super hero
	Super struct {
		UUID         string `json:"uuid"`
		Name         string `json:"name"`
		FullName     string `json:"full-name"`
		Intelligence string `json:"intelligence"`
		Power        string `json:"power"`
		Occupation   string `json:"occupation"`
		Image        string `json:"image"`
	}
)
