package location

type Location struct {
	Name        string
	Description string
	Exits       map[string]*Location
}

type Locations map[string]*Location
