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

func (l *Location) SetExits(exits Locations) {
	l.Exits = exits
}

func (l *Location) AddItem(item *item.Item) {
	l.Items[item.Name] = item
}

type Locations map[string]*Location
