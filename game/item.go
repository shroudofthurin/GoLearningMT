package game

import "fmt"

type Item struct {
	Common
	Openable bool
	Opened   bool
	Takeable bool
}

func (item *Item) Open() {
	if !item.Openable {
		fmt.Println("\nIt seems that this item cannot be opened.\n")
		return
	}

	item.Opened = true
}

func (item *Item) Close() {
	if !item.Openable {
		fmt.Println("\nIt seems that this item cannot be closed.\n")
		return
	}

	item.Opened = false
}

func (item *Item) Info() {
	fmt.Printf("\nThe %v - %v\n", item.Name, item.Description)

	if item.Openable {
		if item.Opened {
			fmt.Printf("The %v is opened.\n", item.Name)
		} else {
			fmt.Printf("The %v is closed.\n", item.Name)
		}
		fmt.Println("\nContains:")
		item.ListInventory()
	}
}

func NewItem(name, description string, openable, takeable bool) *Item {
	item := Item{
		Common{name, description, 0, make(Items)},
		openable,
		false,
		takeable,
	}

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
