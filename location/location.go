package location

type Location struct {
	Name        string
	Description string
	Exits       map[string]string
}

type Locations map[string]*Location
