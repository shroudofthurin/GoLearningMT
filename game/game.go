package game

import (
	"fmt"

	"github.com/shroudofthurin/GoLearningMT/location"
)

type Game struct {
	Location *location.Location
}

func (g Game) Describe() {
	g.DescribeCurrentLocation()
	g.DescribeExits()
}

func (g Game) DescribeCurrentLocation() {
	fmt.Printf("You are in the %v.\n%v\n", g.Location.Name, g.Location.Description)
}

func (g Game) DescribeExits() {
	fmt.Println("Exits:")
	for k, v := range g.Location.Exits {
		fmt.Printf("%s - %s\n", k, v.Name)
	}
}
