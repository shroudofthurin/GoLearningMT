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

func (g *Game) getItem(name string) (*Item, bool) {
	item, ok := g.Location.Inventory[name]

	if ok {
		return item, ok
	}

	item, ok = g.Inventory[name]

	return item, ok
}

func (g *Game) takeFromInventory(name string) (*Item, bool) {
	item := g.Inventory[name]

	delete(g.Inventory, name)

	return item, true
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
	g.DescribeIndividuals()
	g.DescribeItems()
}

func (g *Game) DescribeCurrentLocation() {
	g.Location.Info()
}

func (g *Game) DescribeExits() {
	fmt.Println("\nExits:")
	g.Location.ListExits()
}

func (g *Game) DescribeIndividuals() {
	fmt.Println("\nIndividuals:")
	g.Location.ListIndividuals()
}

func (g *Game) DescribeItems() {
	fmt.Println("\nItems:")
	g.Location.ListInventory()
}

func (g *Game) LookAt(args ...string) {
	item, ok := g.getItem(args[0])

	if ok {
		item.Info()
		return
	}

	character, ok := g.Location.Individuals[args[0]]

	if ok {
		character.Info()
		return
	}

	fmt.Println("\nThat item or person is not available.")
}

func (g *Game) LookIn(args ...string) {
	item, ok := g.getItem(args[0])

	if !ok {
		printItemError()
		return
	}

	item.Contents()
}

func (g *Game) Read(args ...string) {
	item, ok := g.getItem(args[0])

	if !ok {
		printItemError()
		return
	}

	item.Read()
}

func (g *Game) Open(args ...string) {
	item, ok := g.getItem(args[0])

	if !ok {
		printItemError()
		return
	}

	item.Open()

	fmt.Printf("You opened %v.\n", item.Name)
}

func (g *Game) Close(args ...string) {
	item, ok := g.getItem(args[0])

	if !ok {
		printItemError()
		return
	}

	item.Close()

	fmt.Printf("You closed %v.\n", item.Name)
}

func (g *Game) Take(args ...string) {
	_, ok := g.Location.Inventory[args[0]]

	if !ok {
		printItemError()
		return
	}

	item, ok := g.Location.Take(args[0])

	if !ok {
		fmt.Println("\nIt seems that this item cannot be taken.")
		return
	}

	g.Inventory[args[0]] = item

	fmt.Printf("You took %v.\n", item.Name)
}

func (g *Game) TakeFrom(args ...string) {
	items := strings.Split(args[0], "from")

	take := strings.TrimSpace(items[0])
	from := strings.TrimSpace(items[1])

	container, ok := g.getItem(from)

	if !ok {
		fmt.Printf("It seems that %v is not available.\n", from)
		return
	}

	if !container.Openable {
		fmt.Printf("The %v cannot be opened.\n", container.Name)
		return
	} else if !container.Opened {
		fmt.Println("You need open the item to look inside.")
		return
	}

	_, ok = container.Inventory[take]

	if !ok {
		fmt.Printf("It seems that %v is not available.\n", take)
		return
	}

	item, ok := container.Take(take)

	if !ok {
		fmt.Println("\nIt seems that this item cannot be taken.")
		return
	}

	g.Inventory[take] = item

	fmt.Printf("You took %v from %v.\n", item.Name, container.Name)
}

func (g *Game) Drop(args ...string) {
	_, ok := g.Inventory[args[0]]

	if !ok {
		fmt.Printf("It seems that %v is not in your inventory.\n", args[0])
		return
	}

	item, ok := g.takeFromInventory(args[0])

	g.Location.Inventory[args[0]] = item

	fmt.Printf("You dropped %v in %v.\n", item.Name, g.Location.Name)
}

func (g *Game) PutIn(args ...string) {
	items := strings.Split(args[0], " in ")

	put := strings.TrimSpace(items[0])
	inside := strings.TrimSpace(items[1])

	container, ok := g.getItem(inside)

	if !ok {
		fmt.Printf("It seems that %v is not available.\n", inside)
		return
	}

	if !container.Openable {
		fmt.Printf("The %v cannot be opened.\n", container.Name)
		return
	} else if !container.Opened {
		fmt.Println("You need open the item to put things inside.")
		return
	}

	_, ok = g.getItem(put)

	if !ok {
		fmt.Printf("It seems that %v is not available.\n", put)
		return
	}

	item, ok := g.takeFromInventory(put)

	container.Inventory[put] = item

	fmt.Printf("You took %v from your inventory and put it in %v.\n", item.Name, container.Name)
}

func (g *Game) AskFor(args ...string) {
	items := strings.Split(args[0], " for ")
	fmt.Printf("Going to ask %v for %v.\n", items[0], items[1])
}

func (g *Game) Move(to ...string) {
	location, ok := g.Location.Exits[to[0]]

	if !ok {
		fmt.Println("\nYou can't go that direction!!")
		return
	}

	g.Location = location
	g.Describe()
}

func (g *Game) Play() {
	fmt.Println("Let's Hanami!")
	g.Describe()
	fmt.Printf("\n")

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
		fmt.Printf("\n")
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
		command, item := parseLook(commands)
		return command, item
	case "read":
		item := parseItem(commands[1:])
		return "read", item
	case "inventory":
		return "inventory", ""
	case "open":
		item := parseItem(commands[1:])
		return "open", item
	case "close":
		item := parseItem(commands[1:])
		return "close", item
	case "take":
		command, item := parseTake(commands[1:])
		return command, item
	case "drop":
		item := parseItem(commands[1:])
		return "drop", item
	case "put":
		items := parseItem(commands[1:])
		return "put in", items
	case "ask":
		items := parseItem(commands[1:])
		return "ask", items
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

func parseTake(commands []string) (string, string) {
	command := strings.Join(commands, " ")

	contains := strings.Contains(command, "from")
	if contains {
		return "take from", command
	}

	item := parseItem(commands)
	return "take", item
}

func parseLook(commands []string) (string, string) {
	if len(commands) == 1 {
		return "look", ""
	}

	item := parseItem(commands[2:])
	command := strings.Join(commands[:2], " ")

	if command == "look at" {
		return "look at", item
	} else if command == "look in" {
		return "look in", item
	}
	return "look", ""
}

func parseItem(command []string) string {
	item := strings.Join(command, " ")
	return item
}

func printItemError() {
	fmt.Println("\nThat item is not available.")
}
