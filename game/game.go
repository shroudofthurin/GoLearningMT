package game

import (
	"fmt"
	"strings"

	"github.com/peterh/liner"
)

type Game struct {
	Line        *liner.State
	Location    *Location
	CommandList CommandList
	Inventory   Items
}

func New(line *liner.State, location *Location) *Game {
	game := Game{line, location, make(CommandList), make(Items)}
	return &game
}

func (g *Game) SetCommandList(commands CommandList) {
	g.CommandList = commands
}

func (g *Game) Help(args ...string) {
	fmt.Println("You are playing a text adventure game.")
	fmt.Println("You control your character by entering commands")
}

func (g *Game) DescribeInventory(args ...string) {
	fmt.Println("\nCurrent Inventory:")
	g.Inventory.ListItems()
}

func (g *Game) Describe(args ...string) {
	g.DescribeCurrentLocation()
	g.DescribeExits()
	g.DescribeItems()
}

func (g *Game) DescribeCurrentLocation() {
	g.Location.Info()
}

func (g *Game) DescribeExits() {
	fmt.Println("\nExits:")
	g.Location.ListExits()
}

func (g *Game) DescribeItems() {
	fmt.Println("\nItems:")
	g.Location.ListInventory()
}

func (g *Game) LookAt(args ...string) {
	item, ok := g.Location.Inventory[args[0]]

	if !ok {
		fmt.Println("\nThat item is not available in this location.\n")
		return
	}
	item.Info()
}

func (g *Game) Open(args ...string) {
	item, ok := g.Location.Inventory[args[0]]

	if !ok {
		fmt.Println("\nThat item is not available in this location.\n")
		return
	}

	item.Open()
}

func (g *Game) Close(args ...string) {
	item, ok := g.Location.Inventory[args[0]]

	if !ok {
		fmt.Println("\nThat item is not available in this location.\n")
		return
	}

	item.Close()
}

func (g *Game) Move(to ...string) {
	location, ok := g.Location.Exits[to[0]]

	if !ok {
		fmt.Println("\nYou can't go that direction!!\n")
		return
	}

	g.Location = location
	g.Describe()
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
	case "h", "help":
		return "help", ""
	case "go":
		direction := getDirection(commands[1])
		return "go", direction
	case "look":
		if commands[1] == "at" {
			item := getItem(commands[2:])
			return "look at", item
		}
		return "look", ""
	case "inventory":
		return "inventory", ""
	case "open":
		item := getItem(commands[1:])
		return "open", item
	case "close":
		item := getItem(commands[1:])
		return "close", item
	default:
		fmt.Println("Do not understand your command.")
		return "look", ""
	}
}

func getDirection(command string) string {
	direction := strings.ToLower(command)

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

func getItem(command []string) string {
	item := strings.ToLower(strings.Join(command, " "))
	return item
}
