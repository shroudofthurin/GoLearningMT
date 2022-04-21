package game

import (
	"fmt"
	"strings"

	"github.com/peterh/liner"
	"github.com/shroudofthurin/GoLearningMT/game/command"
	"github.com/shroudofthurin/GoLearningMT/location"
)

type Game struct {
	Line        *liner.State
	Location    *location.Location
	CommandList command.CommandList
}

func New(line *liner.State, location *location.Location) *Game {
	game := Game{line, location, make(command.CommandList)}
	return &game
}

func (g *Game) SetCommandList(commands command.CommandList) {
	g.CommandList = commands
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

	fmt.Println("Let's Hanami!\n")
	g.Describe()

	for {
		cmd, err := g.Line.Prompt("What do you want to do? ")

		if err == liner.ErrPromptAborted {
			fmt.Println("Aborted.")
		}

		action, command := parseCommand(cmd)

		if action == "quit" {
			fmt.Println("Quitting game.")
			break
		}

		g.CommandList[action].Action(command)
		g.Describe()
	}
}

func parseCommand(cmd string) (string, string) {
	commands := strings.Split(cmd, " ")

	switch commands[0] {
	case "q", "quit":
		return "quit", "quit"
	case "go":
		direction := getDirection(cmd)
		return "go", direction
	default:
		fmt.Println("Do not understand")
		return "", ""
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
	} else {
		return ""
	}
}
