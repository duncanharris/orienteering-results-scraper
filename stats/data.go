package stats

type (
	CourseData struct {
		Name     string
		Entrants []Entrant
	}

	Entrant struct {
		Name, Club, AgeCat string
	}
)
