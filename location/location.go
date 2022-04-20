package location

import (
	"github.com/shroudofthurin/GoLearningMT/item"
)

type Location struct {
	Name        string
	Description string
	Exits       map[string]*Location
	Items       map[string]*item.Item
}

func New(name, description string) *Location {
	location := Location{
		name,
		description,
		make(map[string]*Location),
		make(map[string]*item.Item),
	}

	return &location
}

func (l *Location) AddItem(item *item.Item) {
	l.Items[item.Name] = item
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
