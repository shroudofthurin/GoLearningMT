package game

import (
	"fmt"
	"strings"

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
	fmt.Println("\nExits:")
	for k, v := range g.Location.Exits {
		fmt.Printf("%s - %s\n", strings.Title(k), v.Name)
	}
	fmt.Println("\n")
}

func (g *Game) Move(to string) {
	location, ok := g.Location.Exits[to]

	if !ok {
		fmt.Println("\nYou can't go that direction!!\n")
		return
	}

	g.Location = location
}

func (g *Game) Play() {
	var err error
	var cmd string

	fmt.Println("Let's Hanami!\n")
	g.Describe()

	for {
		cmd, err = g.Line.Prompt("What do you want to do? ")

		direction := getDirection(cmd)
		fmt.Printf("Direction: %v\n", direction)

		if direction == "quit" {
			fmt.Println("Quitting game.")
			break
		} else if err == liner.ErrPromptAborted {
			fmt.Println("Aborted.")
		}

		g.Move(direction)
		g.Describe()
	}
}

func getDirection(cmd string) string {
	direction := strings.ToLower(cmd)

	if direction == "n" || strings.Contains(direction, "north") {
		return "north"
	} else if direction == "s" || strings.Contains(direction, "south") {
		return "south"
	} else if direction == "e" || strings.Contains(direction, "east") {
		return "east"
	} else if direction == "w" || strings.Contains(direction, "west") {
		return "west"
	} else if direction == "q" || strings.Contains(direction, "quit") {
		return "exit"
	} else {
		return ""
	}
}
