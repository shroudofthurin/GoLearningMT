package location

import (
	"github.com/shroudofthurin/GoLearningMT/item"
)

type Location struct {
	Name        string
	Description string
	Exits       Locations
	Items       item.Items
}

func New(name, description string) *Location {
	location := Location{
		name,
		description,
		make(Locations),
		make(item.Items),
	}

	return &location
}

type Locations map[string]*Location

func SetExits(locations Locations, names ...string) {
	for i := 1; i < len(names); i += 2 {
		locations[names[0]].Exits[names[i]] = locations[names[i+1]]
	}
}

func SetItems(locations Locations, items item.Items, names ...string) {
	for i := 1; i < len(names); i++ {
		locations[names[0]].Items[names[i]] = items[names[i]]
	}
}
