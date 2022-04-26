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
	fmt.Println("You control your character by entering commands.")

	command, ok := g.CommandList[args[0]]

	if ok {
		command.LongDescription()
		return
	}

	for _, command := range g.CommandList {
		command.ShortDescription()
	}
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
		printItemError()
		return
	}
	item.Info()
}

func (g *Game) LookIn(args ...string) {
	item, ok := g.Location.Inventory[args[0]]

	if !ok {
		printItemError()
		return
	}
	item.Contents()
}

func (g *Game) Open(args ...string) {
	item, ok := g.Location.Inventory[args[0]]

	if !ok {
		printItemError()
		return
	}

	item.Open()
}

func (g *Game) Close(args ...string) {
	item, ok := g.Location.Inventory[args[0]]

	if !ok {
		printItemError()
		return
	}

	item.Close()
}

func (g *Game) Take(args ...string) {
	_, ok := g.Location.Inventory[args[0]]

	if !ok {
		printItemError()
		return
	}

	item, ok := g.Location.Take(args[0])

	if !ok {
		fmt.Println("\nIt seems that this item cannot be taken.\n")
		return
	}

	g.Inventory[args[0]] = item
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
	commands := strings.Split(strings.ToLower(cmd), " ")

	switch commands[0] {
	case "q", "quit":
		return "quit", "quit"
	case "h", "help":
		command := ""
		if len(commands) > 1 {
			command = strings.Join(commands[1:], " ")
		}
		return "help", command
	case "go":
		direction := parseDirection(commands[1])
		return "go", direction
	case "look":
		if len(commands) == 1 {
			return "look", ""
		}

		item := parseItem(commands[2:])
		command := strings.Join(commands[:2], " ")

		if command == "look at" {
			return "look at", item
		}
		return "look in", item
	case "inventory":
		return "inventory", ""
	case "open":
		item := parseItem(commands[1:])
		return "open", item
	case "close":
		item := parseItem(commands[1:])
		return "close", item
	case "take":
		item := parseItem(commands[1:])
		return "take", item
	default:
		fmt.Println("Do not understand your command.")
		return "look", ""
	}
}

func parseDirection(command string) string {
	if command == "n" || strings.Contains(command, "north") {
		return "north"
	} else if command == "s" || strings.Contains(command, "south") {
		return "south"
	} else if command == "e" || strings.Contains(command, "east") {
		return "east"
	} else if command == "w" || strings.Contains(command, "west") {
		return "west"
	} else {
		return ""
	}
}

func parseItem(command []string) string {
	item := strings.Join(command, " ")
	return item
}

func printItemError() {
	fmt.Println("\nThat item is not available in this location.\n")
}
