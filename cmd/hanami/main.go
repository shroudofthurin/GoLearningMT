package main

import (
	"fmt"

	"github.com/shroudofthurin/GoLearningMT/location"
)

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
			"north": "porch",
			"south": "park",
		},
	},
	"porch": &location.Location{
		"Porch",
		"You can see the street, your mailbox, and your flower garden.",
		map[string]string{
			"south": "street",
			"west":  "flowergarden",
			"east":  "mailbox",
			"north": "entryway",
		},
	},
	"mailbox": &location.Location{
		"Mailbox",
		"It's where you recieve your mail.",
		map[string]string{
			"west": "porch",
		},
	},
	"flowergarden": &location.Location{
		"Flower Garden",
		"Your flower are in full bloom.",
		map[string]string{
			"east": "porch",
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
			"east":  "livingroom",
			"west":  "bedroom",
		},
	},
	"livingroom": &location.Location{
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
			"south": "cherryblossoms",
		},
	},
	"cherryblossoms": &location.Location{
		"Cherry Blossoms",
		"The cherry blossom trees are in full bloom. It looks like there is space for celebrating Hanami.",
		map[string]string{
			"north": "park",
		},
	},
}

func main() {
	fmt.Println("Let's Hanami!")
	fmt.Println(locations["store"])
}
