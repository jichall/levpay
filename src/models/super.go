package models

type (
	// Super is an abstraction of the used model within the REST API provided by
	// the software. It is used to return data to the client when needed
	Super struct {
		Name         string
		FullName     string
		Intelligence string
		Power        string
		Occupation   string
		Image        string
	}

	// Supers is a collection of Super
	Supers []Super
)
