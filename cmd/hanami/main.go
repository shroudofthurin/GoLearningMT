package main

import (
	"github.com/peterh/liner"
	"github.com/shroudofthurin/GoLearningMT/game"
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

func CreateLocations() location.Locations {
	locations := location.Locations{
		"store": &location.Location{
			"Store",
			"A place where you can buy ice.",
			make(map[string]*location.Location),
		},
		"street": &location.Location{
			"Street",
			"At one end is a store and the other your porch.",
			make(map[string]*location.Location),
		},
		"front yard": &location.Location{
			"Front Yard",
			"It's where you have your garden, mailbox, and leads you to your porch.",
			make(map[string]*location.Location),
		},
		"porch": &location.Location{
			"Porch",
			"You can see the your front yard from here.",
			make(map[string]*location.Location),
		},
		"flower garden": &location.Location{
			"Flower Garden",
			"Your flower are in full bloom.",
			make(map[string]*location.Location),
		},
		"entryway": &location.Location{
			"Entryway",
			"A place to put your bags, keys, and shoes.",
			make(map[string]*location.Location),
		},
		"kitchen": &location.Location{
			"Kitchen",
			"A place where meals are made and enjoyed.",
			make(map[string]*location.Location),
		},
		"hallway": &location.Location{
			"Hallway",
			"A place that connects you to other rooms.",
			make(map[string]*location.Location),
		},
		"living room": &location.Location{
			"Living Room",
			"A place to relax and watch TV.",
			make(map[string]*location.Location),
		},
		"bedroom": &location.Location{
			"Bedroom",
			"A place where you sleep and study.",
			make(map[string]*location.Location),
		},
		"closet": &location.Location{
			"Closet",
			"A place jackets, blankets, and games are stored.",
			make(map[string]*location.Location),
		},
		"engawa": &location.Location{
			"Engawa",
			"A place where you can sit and relax, or follow the gravel path.",
			make(map[string]*location.Location),
		},
		"gravel path": &location.Location{
			"Gravel Path",
			"A gravel path that leads you to your garden.",
			make(map[string]*location.Location),
		},
		"garden": &location.Location{
			"Garden",
			"There is fruit and vegetables a plenty. The strawberries are ready to be picked.",
			make(map[string]*location.Location),
		},
		"park": &location.Location{
			"Park",
			"This lush green park is perfect for celebrating hanami.",
			make(map[string]*location.Location),
		},
		"cherry blossoms": &location.Location{
			"Cherry Blossoms",
			"The cherry blossom trees are in full bloom. It looks like there is space for celebrating Hanami.",
			make(map[string]*location.Location),
		},
	}
	locations["store"].Exits = map[string]*location.Location{
		"east": locations["street"],
	}
	locations["street"].Exits = map[string]*location.Location{
		"north": locations["front yard"],
		"south": locations["park"],
		"west":  locations["store"],
	}
	locations["park"].Exits = map[string]*location.Location{
		"north": locations["street"],
		"south": locations["cherry blossoms"],
	}
	locations["cherry blossoms"].Exits = map[string]*location.Location{
		"north": locations["park"],
	}
	locations["front yard"].Exits = map[string]*location.Location{
		"north": locations["porch"],
		"south": locations["street"],
		"west":  locations["flower garden"],
	}
	locations["flower garden"].Exits = map[string]*location.Location{
		"east": locations["front yard"],
	}
	locations["porch"].Exits = map[string]*location.Location{
		"north": locations["entryway"],
		"south": locations["front yard"],
	}
	locations["entryway"].Exits = map[string]*location.Location{
		"north": locations["hallway"],
		"south": locations["porch"],
		"east":  locations["kitchen"],
	}
	locations["kitchen"].Exits = map[string]*location.Location{
		"west": locations["entryway"],
	}
	locations["hallway"].Exits = map[string]*location.Location{
		"north": locations["engawa"],
		"south": locations["entryway"],
		"east":  locations["living room"],
		"west":  locations["bedroom"],
	}
	locations["living room"].Exits = map[string]*location.Location{
		"west": locations["hallway"],
	}
	locations["bedroom"].Exits = map[string]*location.Location{
		"east":  locations["hallway"],
		"south": locations["closet"],
	}
	locations["closet"].Exits = map[string]*location.Location{
		"north": locations["bedroom"],
	}
	locations["engawa"].Exits = map[string]*location.Location{
		"north": locations["gravel path"],
		"south": locations["hallway"],
	}
	locations["gravel path"].Exits = map[string]*location.Location{
		"north": locations["garden"],
		"south": locations["engawa"],
	}
	locations["garden"].Exits = map[string]*location.Location{
		"south": locations["gravel path"],
	}

	return locations
}

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	locations := CreateLocations()

	game := &game.Game{line, locations["front yard"]}
	game.Play()
}
