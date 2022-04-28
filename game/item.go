package game

import "fmt"

type Item struct {
	Common
	Openable bool
	Opened   bool
	Takeable bool
	Text     string
}

func NewItem(name, description, text string, points int, openable, takeable bool) *Item {
	item := Item{
		Common{name, description, points, make(Items)},
		openable,
		false,
		takeable,
		text,
	}

	return &item
}

func (item *Item) Open() {
	if !item.Openable {
		fmt.Println("\nIt seems that this item cannot be opened.")
		return
	}

	item.Opened = true
}

func (item *Item) Close() {
	if !item.Openable {
		fmt.Println("\nIt seems that this item cannot be closed.")
		return
	}

	item.Opened = false
}

func (item *Item) Read() {
	if len([]rune(item.Text)) == 0 {
		fmt.Println("\nIt seems that this item cannot be read.")
		return
	}

	fmt.Println(item.Text)
}

func (item *Item) Info() {
	fmt.Printf("\nThe %v - %v\n", item.Name, item.Description)

	if item.Openable {
		if item.Opened {
			fmt.Printf("The %v is opened.\n", item.Name)
		} else {
			fmt.Printf("The %v is closed.\n", item.Name)
		}
	}
}

func (item Item) Contents() {
	if !item.Openable {
		fmt.Printf("The %v cannot be opened.\n", item.Name)
		return
	}

	if !item.Opened {
		fmt.Println("You need open the item to look inside.")
		return
	}

	fmt.Println("\nContains:")
	item.ListInventory()
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
