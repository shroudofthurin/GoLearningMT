package game

import (
	"fmt"
	"strings"

	"github.com/peterh/liner"
	"github.com/shroudofthurin/GoLearningMT/location"
)

type Command struct {
	Name     string
	Help     string
	LongHelp string
	Action   func(args ...string)
}

func NewCommand(name, help, longhelp string, action func(args ...string)) *Command {
	command := Command{name, help, longhelp, action}
	return &command
}

type CommandList map[string]*Command

type Game struct {
	Line        *liner.State
	Location    *location.Location
	CommandList CommandList
}

func New(line *liner.State, location *location.Location) *Game {
	game := Game{line, location, make(CommandList)}
	return &game
}

func (g *Game) SetCommandList(commands CommandList) {
	g.CommandList = commands
}

func (g Game) Describe(args ...string) {

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

func (g *Game) Move(to ...string) {
	location, ok := g.Location.Exits[to[0]]

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
	}
}

func parseCommand(cmd string) (string, string) {
	commands := strings.Split(cmd, " ")

	switch commands[0] {
	case "q", "quit":
		return "quit", "quit"
	case "go":
		direction := getDirection(commands[1])
		return "go", direction
	case "describe":
		return "describe", ""
	default:
		fmt.Println("Do not understand")
		return "describe", ""
	}
}

func getDirection(command string) string {
	direction := strings.ToLower(command)

	if direction == "n" || direction == "north" {
		return "north"
	} else if direction == "s" || direction == "south" {
		return "south"
	} else if direction == "e" || direction == "east" {
		return "east"
	} else if direction == "w" || direction == "west" {
		return "west"
	} else {
		return ""
	}
}
