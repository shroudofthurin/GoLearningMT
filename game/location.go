package game

import (
	"fmt"
	"strings"
)

type Location struct {
	Common
	Exits Locations
}

func NewLocation(name, description string) *Location {
	location := Location{
		Common{name, description, 0, make(Items)},
		make(Locations),
	}

	return &location
}

func (location Location) ListExits() {
	for k, v := range location.Exits {
		fmt.Printf("%s - %s\n", strings.Title(k), v.Name)
	}
}

type Locations map[string]*Location

func SetExits(locations Locations, names ...string) {
	for i := 1; i < len(names); i += 2 {
		locations[names[0]].Exits[names[i]] = locations[names[i+1]]
	}
}

func SetLocationInventory(locations Locations, items Items, names ...string) {
	for i := 1; i < len(names); i++ {
		locations[names[0]].Inventory[names[i]] = items[names[i]]
	}
}
