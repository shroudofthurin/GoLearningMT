package main

import (
	"github.com/peterh/liner"
	"github.com/shroudofthurin/GoLearningMT/game"
)

func CreateItems() game.Items {
	items := game.Items{
		"tote bag": game.NewItem(
			"Tote Bag",
			"A large bag suitable for carring lots of things.",
			true,
			true,
		),
		"invitation": game.NewItem(
			"Invitation",
			"A written request inviting you to go somewhere.",
			false,
			true,
		),
		"picnic blanket": game.NewItem(
			"Picnic Blanket",
			"A soft blanket that is perfect for enjoying outdoor activities.",
			false,
			true,
		),
		"cards": game.NewItem(
			"Cards",
			"A standard deck of cards.",
			false,
			true,
		),
		"strawberries": game.NewItem(
			"Strawberries",
			"Small red fruits, which are soft and juicy and ready to be eaten.",
			false,
			true,
		),
		"ice": game.NewItem(
			"Ice",
			"A large, portable bag of ice.",
			false,
			true,
		),
		"cups": game.NewItem(
			"Cups",
			"20 plastic, single use, disposable cups.",
			false,
			true,
		),
		"party checklist": game.NewItem(
			"Party Checklist",
			"A list of items required for the Hanami Party",
			false,
			true,
		),
		"wallet": game.NewItem(
			"Wallet",
			"A pocket-sized flat poketbook for holding money and credit cards.",
			false,
			false,
		),
		"cake": game.NewItem(
			"Cake",
			"A dense, velvety pound cake.",
			false,
			true,
		),
		"fridge": game.NewItem(
			"Fridge",
			"A valuable household appliance, that keeps food and drinks cool.",
			true,
			false,
		),
		"television": game.NewItem(
			"Television",
			"A simple LCD TV mounted on the wall.",
			false,
			false,
		),
		"mailbox": game.NewItem(
			"Mailbox",
			"A letterbox used for receiving incoming mail.",
			true,
			false,
		),
		"flower": game.NewItem(
			"Flower",
			"A beautiful plot of fully bloomed tulips.",
			false,
			false,
		),
		"cherry blossoms": game.NewItem(
			"Cherry Blossoms",
			"A bunch of fully bloomed cheery blossom trees, with beautiful pink and white flowers.",
			false,
			false,
		),
		"cash register": game.NewItem(
			"Cash Register",
			"A drawer for money and totals, displays, and records the amount of each sale.",
			false,
			false,
		),
	}

	game.SetItemInventory(items, "mailbox", "invitation", "party checklist")
	game.SetItemInventory(items, "fridge", "cake")
	game.SetItemInventory(items, "tote bag", "wallet")

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

	game.SetLocationInventory(locations, items, "store", "ice", "cups", "cash register")
	game.SetLocationInventory(locations, items, "garden", "strawberries")
	game.SetLocationInventory(locations, items, "flower garden", "flower")
	game.SetLocationInventory(locations, items, "entryway", "tote bag")
	game.SetLocationInventory(locations, items, "living room", "television")
	game.SetLocationInventory(locations, items, "kitchen", "fridge")
	game.SetLocationInventory(locations, items, "bedroom", "picnic blanket")
	game.SetLocationInventory(locations, items, "closet", "cards")
	game.SetLocationInventory(locations, items, "cherry blossoms", "cherry blossoms")
	game.SetLocationInventory(locations, items, "front yard", "mailbox")

	return locations
}

func CreateCommandList(g *game.Game) game.CommandList {
	commands := game.CommandList{
		"go": game.NewCommand(
			"Go",
			"go <north|south|east|west>: go in the given compass direction.",
			"To move throughout the space, simply type go followed by the direction, e.g \"go north\".",
			g.Move,
		),
		"look": game.NewCommand(
			"Look",
			"look: look around.",
			"To examine your current location, type \"look\". It will describe the current location, exits, and items.",
			g.Describe,
		),
		"look at": game.NewCommand(
			"Look At",
			"look at <item>: look at a specific item.",
			"To exame an item further, type look at followed by the name of the item, e.g. \"look at flowers\".",
			g.LookAt,
		),
		"look in": game.NewCommand(
			"Look In",
			"look in <item>: look inside a specific item.",
			"To exame what is inside an item, type look in followed by the name of the item, e.g. \"look in mailbox\".",
			g.LookIn,
		),
		"inventory": game.NewCommand(
			"Inventory",
			"inventory:  list items in your inventory.",
			"To view your current inventory, type \"inventory\".",
			g.DescribeInventory,
		),
		"open": game.NewCommand(
			"Open",
			"open <item>: open a specific item.",
			"To open an item, type open followed by the name of the item, e.g. \"open mailbox\".",
			g.Open,
		),
		"close": game.NewCommand(
			"Close",
			"close <item>: close a specific item.",
			"To close an item, type close, followed by the name of the item, e.g. \"close mailbox\".",
			g.Close,
		),
		"take": game.NewCommand(
			"Take",
			"take <item>: add an item to your inventory.",
			"To take an item from a location, type take followed by the name of the item, e.g. \"take cake\".",
			g.Take,
		),
		"take from": game.NewCommand(
			"Take From",
			"take <item> from <item>: take an item from another item and add it to your inventory.",
			"To take an item from another item, type take followed by the name of the item from container item, e.g. \"take cake from fridge\".",
			g.TakeFrom,
		),
		"drop": game.NewCommand(
			"Drop",
			"drop <item>: drop an item from your inventory.",
			"To drop an item from your inventory, type drop, followed by the name of the item, e.g. \"drop cake\".",
			g.Drop,
		),
		"put in": game.NewCommand(
			"Put In",
			"put <item> in <item>: take an item from inventory, and put it in another item.",
			"To put an item from your inventory into another item, e.g. \"put wallet in tote bag\".",
			g.PutIn,
		),
		"help": game.NewCommand(
			"Help",
			"help <\"\"|command>: list all commands or a specific one.",
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
