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
			"",
			1,
			true,
			true,
		),
		"invitation": game.NewItem(
			"Invitation",
			"A written request inviting you to go somewhere.",
			"Hanami Picnic!\n"+
				"You are invited to a Hanami Flower Viewing Picnic.\n"+
				"The picnic will take place in the park this afternoon!\n"+
				"We have included a checklist of items needed for the picnic.\n"+
				"Please collect all of the items and meet us at the park!",
			1,
			false,
			true,
		),
		"party checklist": game.NewItem(
			"Party Checklist",
			"A list of items required for the Hanami Party",
			"Hanami Picnic Checklist.\n"+
				"- picnic blanket\n"+
				"- cards\n"+
				"- strawberries\n"+
				"- ice\n"+
				"- cups\n"+
				"- green tea\n"+
				"- cake",
			1,
			false,
			true,
		),
		"green tea": game.NewItem(
			"Green Tea",
			"The perfect spring drink to bring with you to a hanami picnic.",
			"",
			1,
			false,
			true,
		),
		"picnic blanket": game.NewItem(
			"Picnic Blanket",
			"A soft blanket that is perfect for enjoying outdoor activities.",
			"",
			1,
			false,
			true,
		),
		"cards": game.NewItem(
			"Cards",
			"A standard deck of cards.",
			"",
			1,
			false,
			true,
		),
		"strawberries": game.NewItem(
			"Strawberries",
			"Small red fruits, which are soft and juicy and ready to be eaten.",
			"",
			1,
			false,
			true,
		),
		"ice": game.NewItem(
			"Ice",
			"A large, portable bag of ice.",
			"",
			1,
			false,
			true,
		),
		"cups": game.NewItem(
			"Cups",
			"20 plastic, single use, disposable cups.",
			"",
			1,
			false,
			true,
		),
		"wallet": game.NewItem(
			"Wallet",
			"A pocket-sized flat poketbook for holding money and credit cards.",
			"",
			0,
			false,
			false,
		),
		"cake": game.NewItem(
			"Cake",
			"A dense, velvety pound cake.",
			"",
			1,
			false,
			true,
		),
		"fridge": game.NewItem(
			"Fridge",
			"A valuable household appliance, that keeps food and drinks cool.",
			"",
			0,
			true,
			false,
		),
		"hat": game.NewItem(
			"Hat",
			"A hat that will protect you from the sun's UV rays.",
			"",
			0,
			false,
			true,
		),
		"mailbox": game.NewItem(
			"Mailbox",
			"A letterbox used for receiving incoming mail.",
			"",
			0,
			true,
			false,
		),
		"flower": game.NewItem(
			"Flower",
			"A beautiful plot of fully bloomed tulips.",
			"",
			0,
			false,
			true,
		),
		"cherry blossoms": game.NewItem(
			"Cherry Blossoms",
			"A bunch of fully bloomed cheery blossom trees, "+
				"with beautiful pink and white flowers.",
			"",
			0,
			false,
			false,
		),
		"koi pond": game.NewItem(
			"Koi Pond",
			"A relaxing and peaceful koi pond with dozens of beautiful koi fish.",
			"",
			0,
			false,
			false,
		),
	}

	game.SetItemInventory(items, "mailbox", "invitation", "party checklist")
	game.SetItemInventory(items, "fridge", "cake")
	game.SetItemInventory(items, "tote bag", "wallet")

	return items
}

func CreateCharacters(items game.Items) game.Characters {
	characters := game.Characters{
		"store clerk": game.NewCharacter(
			"Store Clerk",
			"Someone who works at the store and might have some items for you.",
			"Hello! Welcome to the Store, where everything is free!",
			2,
			true,
			false,
		),
		"lazy cat": game.NewCharacter(
			"Lazy Cat",
			"A fluffy cat is napping the afternoon away in the morning sun.",
			"nyaaaaaaoooooo. gorogoro.",
			0,
			false,
			false,
		),
		"obaa-chan": game.NewCharacter(
			"Obaa-chan",
			"A sweet old lady who has been running a small tea house "+
				"her whole life. She has the best green tea in town.",
			"Good afternoon. We are giving away green tea for free to "+
				"celebrate the hanami season. There is limited supply.",
			1,
			true,
			false,
		),
		"ren": game.NewCharacter(
			"Ren",
			"Your friend who is organizing the hanami viewing picnic.",
			"I can't wait for everyone to arrive with food and drinks!!",
			1,
			false,
			true,
		),
	}

	game.SetCharacterInventory(characters, items, "store clerk", "ice", "cups")
	game.SetCharacterInventory(characters, items, "obaa-chan", "green tea")

	return characters
}

func CreateLocations(items game.Items, characters game.Characters) game.Locations {
	locations := game.Locations{
		"store": game.NewLocation(
			"Store",
			"A place where you can buy ice.",
			0,
		),
		"street": game.NewLocation(
			"Street",
			"At one end is a store and the other your porch.",
			0,
		),
		"front yard": game.NewLocation(
			"Front Yard",
			"It's where you have your mailbox.",
			0,
		),
		"porch": game.NewLocation(
			"Porch",
			"You can see the your front yard from here.",
			0,
		),
		"flower garden": game.NewLocation(
			"Flower Garden",
			"Your flower are in full bloom.",
			0,
		),
		"entryway": game.NewLocation(
			"Entryway",
			"A place to put your bags, keys, and shoes.",
			0,
		),
		"kitchen": game.NewLocation(
			"Kitchen",
			"A place where meals are made and enjoyed.",
			0,
		),
		"hallway": game.NewLocation(
			"Hallway",
			"A place that connects you to other rooms.",
			0,
		),
		"living room": game.NewLocation(
			"Living Room",
			"A place to relax and watch TV.",
			0,
		),
		"bedroom": game.NewLocation(
			"Bedroom",
			"A place where you sleep and study.",
			0,
		),
		"tea house": game.NewLocation(
			"Tea House",
			"A place to that provides quality tea and amazing views.",
			0,
		),
		"engawa": game.NewLocation(
			"Engawa",
			"A place where you can sit and relax, or follow the gravel path.",
			1,
		),
		"gravel path": game.NewLocation(
			"Gravel Path",
			"A gravel path that leads you to your garden.",
			0,
		),
		"garden": game.NewLocation(
			"Garden",
			"There is fruit and vegetables a plenty. The strawberries are ready to be picked.",
			0,
		),
		"park": game.NewLocation(
			"Park",
			"This lush green park is perfect for celebrating hanami.",
			0,
		),
		"cherry blossoms": game.NewLocation(
			"Cherry Blossoms",
			"The cherry blossom trees are in full bloom. It looks like there is space for celebrating Hanami.",
			1,
		),
	}
	game.SetExits(locations, "store", "east", "street")
	game.SetExits(locations, "street", "north", "front yard", "south", "park", "west", "store")
	game.SetExits(locations, "park", "north", "street", "south", "cherry blossoms", "east", "tea house")
	game.SetExits(locations, "cherry blossoms", "north", "park")
	game.SetExits(locations, "tea house", "west", "park")
	game.SetExits(locations, "front yard", "north", "porch", "south", "street", "west", "flower garden")
	game.SetExits(locations, "flower garden", "east", "front yard")
	game.SetExits(locations, "porch", "north", "entryway", "south", "front yard")
	game.SetExits(locations, "entryway", "north", "hallway", "south", "porch", "east", "kitchen")
	game.SetExits(locations, "kitchen", "west", "entryway")
	game.SetExits(locations, "hallway", "north", "engawa", "south", "entryway", "east", "living room", "west", "bedroom")
	game.SetExits(locations, "living room", "west", "hallway")
	game.SetExits(locations, "bedroom", "east", "hallway")
	game.SetExits(locations, "engawa", "north", "gravel path", "south", "hallway")
	game.SetExits(locations, "gravel path", "north", "garden", "south", "engawa")
	game.SetExits(locations, "garden", "south", "gravel path")

	game.SetCharacters(locations, characters, "store", "store clerk")
	game.SetCharacters(locations, characters, "garden", "lazy cat")
	game.SetCharacters(locations, characters, "tea house", "obaa-chan")
	game.SetCharacters(locations, characters, "park", "ren")

	game.SetLocationInventory(locations, items, "garden", "strawberries")
	game.SetLocationInventory(locations, items, "engawa", "koi pond")
	game.SetLocationInventory(locations, items, "flower garden", "flower")
	game.SetLocationInventory(locations, items, "entryway", "tote bag")
	game.SetLocationInventory(locations, items, "living room", "hat")
	game.SetLocationInventory(locations, items, "kitchen", "fridge")
	game.SetLocationInventory(locations, items, "bedroom", "picnic blanket", "cards")
	game.SetLocationInventory(locations, items, "cherry blossoms", "cherry blossoms")
	game.SetLocationInventory(locations, items, "front yard", "mailbox")

	return locations
}

func CreateCommandList(g *game.Game) game.CommandList {
	commands := game.CommandList{
		"go": game.NewCommand(
			"Go",
			"go <north|south|east|west>: go in the given compass direction.",
			"To move throughout the space, simply type go "+
				"followed by the direction, e.g \"go north\".",
			g.Move,
		),
		"look": game.NewCommand(
			"Look",
			"look: look around.",
			"To examine your current location, type \"look\". "+
				"It will describe the current location, exits, and items.",
			g.Describe,
		),
		"look at": game.NewCommand(
			"Look At",
			"look at <item>: look at a specific item.",
			"To exame an item further, type look at followed by the "+
				"name of the item, e.g. \"look at flowers\".",
			g.LookAt,
		),
		"look in": game.NewCommand(
			"Look In",
			"look in <item>: look inside a specific item.",
			"To exame what is inside an item, type look in followed by "+
				"the name of the item, e.g. \"look in mailbox\".",
			g.LookIn,
		),
		"read": game.NewCommand(
			"Read",
			"read <item>: read the text of a specific item.",
			"To read the text of an item, type read followed by "+
				"the name of the item, e.g. \"read invitation\".",
			g.Read,
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
			"To open an item, type open followed by the name of the item, "+
				"e.g. \"open mailbox\".",
			g.Open,
		),
		"close": game.NewCommand(
			"Close",
			"close <item>: close a specific item.",
			"To close an item, type close, followed by the name of the item, "+
				"e.g. \"close mailbox\".",
			g.Close,
		),
		"take": game.NewCommand(
			"Take",
			"take <item>: add an item to your inventory.",
			"To take an item from a location, type take followed by the "+
				"name of the item, e.g. \"take cake\".",
			g.Take,
		),
		"take from": game.NewCommand(
			"Take From",
			"take <item> from <item>: take an item from another item and "+
				"add it to your inventory.",
			"To take an item from another item, type take followed by the "+
				"name of the item from container item, "+
				"e.g. \"take cake from fridge\".",
			g.TakeFrom,
		),
		"drop": game.NewCommand(
			"Drop",
			"drop <item>: drop an item from your inventory.",
			"To drop an item from your inventory, type drop, followed by the "+
				"name of the item, e.g. \"drop cake\".",
			g.Drop,
		),
		"put in": game.NewCommand(
			"Put In",
			"put <item> in <item>: take an item from inventory, "+
				"and put it in another item.",
			"To put an item from your inventory into another item, "+
				"e.g. \"put wallet in tote bag\".",
			g.PutIn,
		),
		"ask": game.NewCommand(
			"Ask",
			"ask <character> for <item>: ask character for item "+
				"in their inventory.",
			"To ask a character for an item in their inventory "+
				"e.g. \"ask store clerk for ice\".",
			g.AskFor,
		),
		"give": game.NewCommand(
			"Give",
			"give <item> to <character>: give an item to a character.",
			"You can give a character an item from your inventory. If you "+
				"give a character the right item, you might get a point or two "+
				"in return.",
			g.Give,
		),
		"say hello": game.NewCommand(
			"Say Hello",
			"say hello to <character>: converse with another character.",
			"To hear what another character has to say, simply type, say "+
				"hello to followed by their name e.g. \"say hello to ren\".",
			g.SayHello,
		),
		"see score": game.NewCommand(
			"See Score",
			"see score: look at your current score.",
			"To see your current score in the game, just type \"see score\".",
			g.SeeScore,
		),
		"help": game.NewCommand(
			"Help",
			"help <\"\"|command>: list all commands or a specific one.",
			"I will help you naviate this game and the commands available",
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
	characters := CreateCharacters(items)
	locations := CreateLocations(items, characters)

	hanamiGame := game.New(line, locations["front yard"])
	commands := CreateCommandList(hanamiGame)
	hanamiGame.SetCommandList(commands)

	hanamiGame.Play()
}
