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
	Points      int
}

func New(line *liner.State, location *Location) *Game {
	game := Game{line, location, make(CommandList), make(Items), 0}
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

func (g *Game) SeeScore(args ...string) {
	fmt.Println("Current Score:", g.Points)
}

func (g *Game) DescribeInventory(args ...string) {
	fmt.Println("Current Inventory:")

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

	fmt.Println("That item or person is not available.")
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
		fmt.Println("It seems that this item cannot be taken.")
		return
	}

	if item.Points > 0 {
		g.Points += item.Points
		item.Points = 0
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

	if item.Points > 0 {
		g.Points += item.Points
		item.Points = 0
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

	character := strings.TrimSpace(items[0])
	want := strings.TrimSpace(items[1])

	individual, ok := g.Location.Individuals[character]

	if !ok {
		fmt.Printf("It seems that %v is not available.\n", character)
		return
	}

	if !individual.Askable {
		fmt.Printf("%v does not have any item to give.\n", individual.Name)
		return
	}

	_, ok = individual.Inventory[want]

	if !ok {
		fmt.Printf("It seems that %v is not available.\n", want)
		return
	}

	item, ok := individual.Take(want)

	if item.Points > 0 {
		g.Points += item.Points
		item.Points = 0
	}

	if !individual.Askable {
		g.Points += individual.Points
		individual.Points = 0
	}

	g.Inventory[want] = item

	fmt.Printf("%v gave you %v.\n", individual.Name, item.Name)
}

func (g *Game) SayHello(args ...string) {
	character, ok := g.Location.Individuals[args[0]]

	if !ok {
		fmt.Printf("It seems that %v is not available.\n", args[0])
		return
	}
	fmt.Println(character.Phrase)
}

func (g *Game) Move(to ...string) {
	location, ok := g.Location.Exits[to[0]]

	if !ok {
		fmt.Println("\nYou can't go that direction!!")
		return
	}

	if location.Points > 0 {
		g.Points += location.Points
		location.Points = 0
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

		if g.Points == 16 {
			fmt.Println("You have collected all of the required items,")
			fmt.Println("talked to the required characters, and")
			fmt.Println("visited the required locations.")
			fmt.Printf("\n\n")
			fmt.Println("You spent the afternoon in the park with all of")
			fmt.Println("your friends, laughing, eating, drinking, and")
			fmt.Println("admiring the beautiful cherry blossoms.")
			fmt.Printf("\n\n")

			break
		}
	}
	fmt.Println("The game has ended!")
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
	case "say":
		command := strings.Join(commands[:3], " ")
		if command != "say hello to" {
			return "help", "say hello"
		}
		items := parseItem(commands[3:])
		return "say hello", items
	case "see":
		command := strings.Join(commands[:2], " ")
		if command != "see score" {
			return "help", "see score"
		}
		return "see score", ""
	default:
		fmt.Println("Do not understand your command.")
		return "help", ""
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
