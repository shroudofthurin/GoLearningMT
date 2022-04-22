package game

import "fmt"

type Item struct {
	Name     string
	Openable bool
	Contains Items
}

func NewItem(name string, openable bool) *Item {
	item := Item{name, openable, make(Items)}

	return &item
}

func SetContains(items Items, names ...string) {
	for i := 1; i < len(names); i++ {
		items[names[0]].Contains[names[i]] = items[names[i]]
	}
}

type Items map[string]*Item

func (items Items) ListItems() {
	if len(items) == 0 {
		fmt.Println("{}")
	} else {
		for _, v := range items {
			fmt.Printf("%s\n", v.Name)
		}
	}
}
