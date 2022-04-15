package main

import (
	"fmt"

	"github.com/peterh/liner"
	"github.com/shroudofthurin/GoLearningMT/item"
	"github.com/shroudofthurin/GoLearningMT/location"
)

var items = map[string]*item.Item{
	"tote bag":      {"Tote Bag"},
	"invitation":    {"Invitation"},
	"blanket":       {"Blanket"},
	"cards":         {"Cards"},
	"strawberries":  {"Strawberries"},
	"ice":           {"Ice"},
	"cups":          {"Cups"},
	"scissor":       {"Scissor"},
	"wallet":        {"Wallet"},
	"cake":          {"Cake"},
	"refrigerator":  {"Refrigerator"},
	"tv":            {"Television"},
	"mailbox":       {"Mailbox"},
	"flower":        {"Flower"},
	"sakura":        {"Sakura Tree"},
	"cash register": {"Cash Register"},
}

var locations = location.Locations{
	"store": &location.Location{
		"Store",
		"A place where you can buy ice.",
		map[string]string{
			"east": "street",
		},
	},
	"street": &location.Location{
		"Street",
		"At one end is a store and the other your porch.",
		map[string]string{
			"west":  "store",
			"north": "front yard",
			"south": "park",
		},
	},
	"front yard": &location.Location{
		"Front Yard",
		"It's where you have your garden, mailbox, and leads you to your porch.",
		map[string]string{
			"north": "porch",
			"south": "street",
			"west":  "flower garden",
		},
	},
	"porch": &location.Location{
		"Porch",
		"You can see the your front yard from here.",
		map[string]string{
			"south": "front yard",
			"north": "entryway",
		},
	},
	"flower garden": &location.Location{
		"Flower Garden",
		"Your flower are in full bloom.",
		map[string]string{
			"east": "front yard",
		},
	},
	"entrway": &location.Location{
		"Entryway",
		"A place to put your bags, keys, and shoes.",
		map[string]string{
			"north": "hallway",
			"south": "porch",
			"east":  "kitchen",
		},
	},
	"kitchen": &location.Location{
		"Kitchen",
		"A place where meals are made and enjoyed.",
		map[string]string{
			"west": "entryway",
		},
	},
	"hallway": &location.Location{
		"Hallway",
		"A place that connects you to other rooms.",
		map[string]string{
			"north": "engawa",
			"south": "entryway",
			"east":  "living room",
			"west":  "bedroom",
		},
	},
	"living room": &location.Location{
		"Living Room",
		"A place to relax and watch TV.",
		map[string]string{
			"west": "hallway",
		},
	},
	"bedroom": &location.Location{
		"Bedroom",
		"A place where you sleep and study.",
		map[string]string{
			"east":  "hallway",
			"south": "closet",
		},
	},
	"closet": &location.Location{
		"Closet",
		"A place jackets, blankets, and games are stored.",
		map[string]string{
			"north": "bedroom",
		},
	},
	"engawa": &location.Location{
		"Engawa",
		"A place where you can sit and relax, or follow the gravel path.",
		map[string]string{
			"south": "hallway",
			"north": "path",
		},
	},
	"path": &location.Location{
		"Path",
		"A path that leads you to your garden.",
		map[string]string{
			"south": "engawa",
			"north": "garden",
		},
	},
	"garden": &location.Location{
		"Garden",
		"There is fruit and vegetables a plenty. The strawberries are ready to be picked.",
		map[string]string{
			"south": "path",
		},
	},
	"park": &location.Location{
		"Park",
		"This lush green park is perfect for celebrating hanami.",
		map[string]string{
			"north": "street",
			"south": "cherry blossoms",
		},
	},
	"cherry blossoms": &location.Location{
		"Cherry Blossoms",
		"The cherry blossom trees are in full bloom. It looks like there is space for celebrating Hanami.",
		map[string]string{
			"north": "park",
		},
	},
}

func main() {
	line := liner.NewLiner()
	defer line.Close()

	fmt.Println("Let's Hanami!")
}
