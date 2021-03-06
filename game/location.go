package game

import (
	"fmt"
	"strings"
)

type Location struct {
	Common
	Exits       Locations
	Individuals Characters
}

func NewLocation(name, description string, points int) *Location {
	location := Location{
		Common{name, description, points, make(Items)},
		make(Locations),
		make(Characters),
	}

	return &location
}

func (location Location) ListExits() {
	for k, v := range location.Exits {
		fmt.Printf("%s - %s\n", strings.Title(k), v.Name)
	}
}

func (location Location) ListIndividuals() {
	if len(location.Individuals) == 0 {
		fmt.Println("There is no one here.")
		return
	}

	for _, v := range location.Individuals {
		fmt.Printf("%s.\n", v.Name)
	}
}

type Locations map[string]*Location

func SetExits(locations Locations, names ...string) {
	for i := 1; i < len(names); i += 2 {
		locations[names[0]].Exits[names[i]] = locations[names[i+1]]
	}
}

func SetCharacters(locations Locations, characters Characters, names ...string) {
	for i := 1; i < len(names); i++ {
		locations[names[0]].Individuals[names[i]] = characters[names[i]]
	}
}

func SetLocationInventory(locations Locations, items Items, names ...string) {
	for i := 1; i < len(names); i++ {
		locations[names[0]].Inventory[names[i]] = items[names[i]]
	}
}
