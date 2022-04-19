package game

import (
	"fmt"

	"github.com/peterh/liner"
	"github.com/shroudofthurin/GoLearningMT/location"
)

type Game struct {
	Line     *liner.State
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

func (g *Game) Play() {
	fmt.Println("Let's Hanami!")

	if cmd, err := g.Line.Prompt("What do you want to do? "); err == nil {
		fmt.Print("Got: ", cmd)
	} else if err == liner.ErrPromptAborted {
		fmt.Println("Aborted.")
	} else {
		fmt.Println("Error reading line: ", err)
	}
}
