package game

import "fmt"

type Item struct {
	Common
	Openable bool
}

func (item *Item) Open() {
	if !item.Openable {
		fmt.Println("\nIt seems that this item cannot be opened.\n")
		return
	}
	fmt.Println("Container Items: ")
	item.ListInventory()
}

func NewItem(name, description string, openable bool) *Item {
	item := Item{Common{name, description, 0, make(Items)}, openable}

	return &item
}

func SetItemInventory(items Items, names ...string) {
	for i := 1; i < len(names); i++ {
		items[names[0]].Inventory[names[i]] = items[names[i]]
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
