package superapi

type (
	// Response represents a response from SuperAPI with the needed fields.
	Response struct {
		ID string `json:"id"`
		Name string `json:"name"`

		Image Image `json:"image"`
		Powerstats Powerstats `json:"powerstats"`
		Connections Connections `json:"connections"`
	}

	// Powerstats represents the stats of a hero
	Powerstats struct {
		Intelligence string `json:"intelligence"`
		Strength string `json:"strength"`
		Speed string `json:"speed"`
		Durability string `json:"durability"`
		Power string `json:"powerstats"`
		Combat string `json:"combat"`
	}

	// Connections represents the connections that a hero has
	Connections struct {
		Affiliation string `json:"affiliation"`
		Relatives string `json:"relatives"`
	}

	// Image is, of course, the image of the hero
	Image struct {
		URL string `json:"url"`
	}
)