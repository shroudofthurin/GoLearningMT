package main

import (
	"github.com/peterh/liner"
	"github.com/shroudofthurin/GoLearningMT/game"
)

func CreateItems() game.Items {
	items := game.Items{
		"tote bag":      game.NewItem("Tote Bag", true),
		"invitation":    game.NewItem("Invitation", false),
		"blanket":       game.NewItem("Blanket", false),
		"cards":         game.NewItem("Cards", false),
		"strawberries":  game.NewItem("Strawberries", false),
		"ice":           game.NewItem("Ice", false),
		"cups":          game.NewItem("Cups", false),
		"checklist":     game.NewItem("Hanami Party Checklist", false),
		"wallet":        game.NewItem("Wallet", false),
		"cake":          game.NewItem("Cake", false),
		"refrigerator":  game.NewItem("Refrigerator", true),
		"tv":            game.NewItem("Television", false),
		"mailbox":       game.NewItem("Mailbox", true),
		"flower":        game.NewItem("Flower", false),
		"sakura":        game.NewItem("Sakura Tree", false),
		"cash register": game.NewItem("Cash Register", false),
	}

	game.SetContains(items, "mailbox", "invitation", "checklist")
	game.SetContains(items, "refrigerator", "cake")
	game.SetContains(items, "tote bag", "wallet")

	return items
}

func CreateLocations(items game.Items) game.Locations {
	locations := game.Locations{
		"store": game.NewLocation(
			"Store",
			"A place where you can buy ice.",
		),
		"street": game.NewLocation(
			"Street",
			"At one end is a store and the other your porch.",
		),
		"front yard": game.NewLocation(
			"Front Yard",
			"It's where you have your mailbox.",
		),
		"porch": game.NewLocation(
			"Porch",
			"You can see the your front yard from here.",
		),
		"flower garden": game.NewLocation(
			"Flower Garden",
			"Your flower are in full bloom.",
		),
		"entryway": game.NewLocation(
			"Entryway",
			"A place to put your bags, keys, and shoes.",
		),
		"kitchen": game.NewLocation(
			"Kitchen",
			"A place where meals are made and enjoyed.",
		),
		"hallway": game.NewLocation(
			"Hallway",
			"A place that connects you to other rooms.",
		),
		"living room": game.NewLocation(
			"Living Room",
			"A place to relax and watch TV.",
		),
		"bedroom": game.NewLocation(
			"Bedroom",
			"A place where you sleep and study.",
		),
		"closet": game.NewLocation(
			"Closet",
			"A place jackets, blankets, and games are stored.",
		),
		"engawa": game.NewLocation(
			"Engawa",
			"A place where you can sit and relax, or follow the gravel path.",
		),
		"gravel path": game.NewLocation(
			"Gravel Path",
			"A gravel path that leads you to your garden.",
		),
		"garden": game.NewLocation(
			"Garden",
			"There is fruit and vegetables a plenty. The strawberries are ready to be picked.",
		),
		"park": game.NewLocation(
			"Park",
			"This lush green park is perfect for celebrating hanami.",
		),
		"cherry blossoms": game.NewLocation(
			"Cherry Blossoms",
			"The cherry blossom trees are in full bloom. It looks like there is space for celebrating Hanami.",
		),
	}
	game.SetExits(locations, "store", "east", "street")
	game.SetExits(locations, "street", "north", "front yard", "south", "park", "west", "store")
	game.SetExits(locations, "park", "north", "street", "south", "cherry blossoms")
	game.SetExits(locations, "cherry blossoms", "north", "park")
	game.SetExits(locations, "front yard", "north", "porch", "south", "street", "west", "flower garden")
	game.SetExits(locations, "flower garden", "east", "front yard")
	game.SetExits(locations, "porch", "north", "entryway", "south", "front yard")
	game.SetExits(locations, "entryway", "north", "hallway", "south", "porch", "east", "kitchen")
	game.SetExits(locations, "kitchen", "west", "entryway")
	game.SetExits(locations, "hallway", "north", "engawa", "south", "entryway", "east", "living room", "west", "bedroom")
	game.SetExits(locations, "living room", "west", "hallway")
	game.SetExits(locations, "bedroom", "south", "closet", "east", "hallway")
	game.SetExits(locations, "closet", "north", "bedroom")
	game.SetExits(locations, "engawa", "north", "gravel path", "south", "hallway")
	game.SetExits(locations, "gravel path", "north", "garden", "south", "engawa")
	game.SetExits(locations, "garden", "south", "gravel path")

	game.SetItems(locations, items, "store", "ice", "cups", "cash register")
	game.SetItems(locations, items, "garden", "strawberries")
	game.SetItems(locations, items, "flower garden", "flower")
	game.SetItems(locations, items, "entryway", "tote bag")
	game.SetItems(locations, items, "living room", "tv")
	game.SetItems(locations, items, "kitchen", "refrigerator")
	game.SetItems(locations, items, "bedroom", "blanket")
	game.SetItems(locations, items, "closet", "cards")
	game.SetItems(locations, items, "cherry blossoms", "sakura")
	game.SetItems(locations, items, "front yard", "mailbox")

	return locations
}

func CreateCommandList(g *game.Game) game.CommandList {
	commands := game.CommandList{
		"go": game.NewCommand(
			"go",
			"go <north|south|east|west>",
			"To move throughout the space, simply type go followed by the direction, e.g \"go north\".",
			g.Move,
		),
		"describe": game.NewCommand(
			"describe",
			"describe",
			"To examine your current location, type \"describe\". It will describe the current location, exits, and items",
			g.Describe,
		),
		"inventory": game.NewCommand(
			"inventory",
			"inventory",
			"To view your current inventory, type \"inventory\".",
			g.DescribeInventory,
		),
		"help": game.NewCommand(
			"help",
			"help <\"\"|command>",
			"I will eventually help you naviate this game and the commands available",
			g.Help,
		),
	}

	return commands
}

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	items := CreateItems()
	locations := CreateLocations(items)

	hanamiGame := game.New(line, locations["front yard"])
	commands := CreateCommandList(hanamiGame)
	hanamiGame.SetCommandList(commands)

	hanamiGame.Play()
}
