package main

import (
	"fmt"

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
	locations["store"].SetExits(location.Locations{
		"east": locations["street"],
	})
	locations["street"].SetExits(location.Locations{
		"north": locations["front yard"],
		"south": locations["park"],
		"west":  locations["store"],
	})
	locations["park"].SetExits(location.Locations{
		"north": locations["street"],
		"south": locations["cherry blossoms"],
	})
	locations["cherry blossoms"].SetExits(location.Locations{
		"north": locations["park"],
	})
	locations["front yard"].SetExits(location.Locations{
		"north": locations["porch"],
		"south": locations["street"],
		"west":  locations["flower garden"],
	})
	locations["flower garden"].SetExits(location.Locations{
		"east": locations["front yard"],
	})
	locations["porch"].SetExits(location.Locations{
		"north": locations["entryway"],
		"south": locations["front yard"],
	})
	locations["entryway"].SetExits(location.Locations{
		"north": locations["hallway"],
		"south": locations["porch"],
		"east":  locations["kitchen"],
	})
	locations["kitchen"].SetExits(location.Locations{
		"west": locations["entryway"],
	})
	locations["hallway"].SetExits(location.Locations{
		"north": locations["engawa"],
		"south": locations["entryway"],
		"east":  locations["living room"],
		"west":  locations["bedroom"],
	})
	locations["living room"].SetExits(location.Locations{
		"west": locations["hallway"],
	})
	locations["bedroom"].SetExits(location.Locations{
		"east":  locations["hallway"],
		"south": locations["closet"],
	})
	locations["closet"].SetExits(location.Locations{
		"north": locations["bedroom"],
	})
	locations["engawa"].SetExits(location.Locations{
		"north": locations["gravel path"],
		"south": locations["hallway"],
	})
	locations["gravel path"].SetExits(location.Locations{
		"north": locations["garden"],
		"south": locations["engawa"],
	})
	locations["garden"].SetExits(location.Locations{
		"south": locations["gravel path"],
	})

	locations["store"].AddItem(items["ice"])
	locations["store"].AddItem(items["cups"])
	locations["store"].AddItem(items["cash register"])
	locations["garden"].AddItem(items["scissor"])
	locations["garden"].AddItem(items["strawberries"])
	locations["garden"].AddItem(items["flower"])
	locations["entryway"].AddItem(items["tote bag"])
	locations["entryway"].AddItem(items["wallet"])
	locations["living room"].AddItem(items["tv"])
	locations["kitchen"].AddItem(items["refrigerator"])
	locations["bedroom"].AddItem(items["blanket"])
	locations["closet"].AddItem(items["cards"])
	locations["cherry blossoms"].AddItem(items["sakura"])
	locations["front yard"].AddItem(items["mailbox"])

	return locations
}

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	items := CreateItems()
	fmt.Println(items["mailbox"])
	locations := CreateLocations(items)

	game := &game.Game{line, locations["front yard"]}
	game.Play()
}
