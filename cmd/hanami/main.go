package main

import (
	"github.com/peterh/liner"
	"github.com/shroudofthurin/GoLearningMT/game"
	"github.com/shroudofthurin/GoLearningMT/item"
	"github.com/shroudofthurin/GoLearningMT/location"
)

func CreateItems() item.Items {
	items := item.Items{
		"tote bag":      {"Tote Bag", true},
		"invitation":    {"Invitation", false},
		"blanket":       {"Blanket", false},
		"cards":         {"Cards", false},
		"strawberries":  {"Strawberries", false},
		"ice":           {"Ice", false},
		"cups":          {"Cups", false},
		"scissor":       {"Scissor", false},
		"wallet":        {"Wallet", false},
		"cake":          {"Cake", false},
		"refrigerator":  {"Refrigerator", true},
		"tv":            {"Television", false},
		"mailbox":       {"Mailbox", true},
		"flower":        {"Flower", false},
		"sakura":        {"Sakura Tree", false},
		"cash register": {"Cash Register", false},
	}

	return items
}

func CreateLocations(items item.Items) location.Locations {
	locations := location.Locations{
		"store": location.New(
			"Store",
			"A place where you can buy ice.",
		),
		"street": location.New(
			"Street",
			"At one end is a store and the other your porch.",
		),
		"front yard": location.New(
			"Front Yard",
			"It's where you have your mailbox.",
		),
		"porch": location.New(
			"Porch",
			"You can see the your front yard from here.",
		),
		"flower garden": location.New(
			"Flower Garden",
			"Your flower are in full bloom.",
		),
		"entryway": location.New(
			"Entryway",
			"A place to put your bags, keys, and shoes.",
		),
		"kitchen": location.New(
			"Kitchen",
			"A place where meals are made and enjoyed.",
		),
		"hallway": location.New(
			"Hallway",
			"A place that connects you to other rooms.",
		),
		"living room": location.New(
			"Living Room",
			"A place to relax and watch TV.",
		),
		"bedroom": location.New(
			"Bedroom",
			"A place where you sleep and study.",
		),
		"closet": location.New(
			"Closet",
			"A place jackets, blankets, and games are stored.",
		),
		"engawa": location.New(
			"Engawa",
			"A place where you can sit and relax, or follow the gravel path.",
		),
		"gravel path": location.New(
			"Gravel Path",
			"A gravel path that leads you to your garden.",
		),
		"garden": location.New(
			"Garden",
			"There is fruit and vegetables a plenty. The strawberries are ready to be picked.",
		),
		"park": location.New(
			"Park",
			"This lush green park is perfect for celebrating hanami.",
		),
		"cherry blossoms": location.New(
			"Cherry Blossoms",
			"The cherry blossom trees are in full bloom. It looks like there is space for celebrating Hanami.",
		),
	}
	location.SetExits(locations, "store", "east", "street")
	location.SetExits(locations, "street", "north", "front yard", "south", "park", "west", "store")
	location.SetExits(locations, "park", "north", "street", "south", "cherry blossoms")
	location.SetExits(locations, "cherry blossoms", "north", "park")
	location.SetExits(locations, "front yard", "north", "porch", "south", "street", "west", "flower garden")
	location.SetExits(locations, "flower garden", "east", "front yard")
	location.SetExits(locations, "porch", "north", "entryway", "south", "front yard")
	location.SetExits(locations, "entryway", "north", "hallway", "south", "porch", "east", "kitchen")
	location.SetExits(locations, "kitchen", "west", "entryway")
	location.SetExits(locations, "hallway", "north", "engawa", "south", "entryway", "east", "living room", "west", "bedroom")
	location.SetExits(locations, "living room", "west", "hallway")
	location.SetExits(locations, "bedroom", "south", "closet", "east", "hallway")
	location.SetExits(locations, "closet", "north", "bedroom")
	location.SetExits(locations, "engawa", "north", "gravel path", "south", "hallway")
	location.SetExits(locations, "gravel path", "north", "garden", "south", "engawa")
	location.SetExits(locations, "garden", "south", "gravel path")

	location.SetItems(locations, items, "store", "ice", "cups", "cash register")
	location.SetItems(locations, items, "garden", "scissor", "strawberries")
	location.SetItems(locations, items, "flower garden", "flower")
	location.SetItems(locations, items, "entryway", "tote bag", "wallet")
	location.SetItems(locations, items, "living room", "tv")
	location.SetItems(locations, items, "kitchen", "refrigerator")
	location.SetItems(locations, items, "bedroom", "blanket")
	location.SetItems(locations, items, "closet", "cards")
	location.SetItems(locations, items, "cherry blossoms", "sakura")
	location.SetItems(locations, items, "front yard", "mailbox")

	return locations
}

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	items := CreateItems()
	locations := CreateLocations(items)

	game := &game.Game{line, locations["front yard"]}
	game.Play()
}
