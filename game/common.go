package game

import "fmt"

type Common struct {
	Name        string
	Description string
	Points      int
	Inventory   Items
}

func (common *Common) Take(name string) (*Item, bool) {
	item := common.Inventory[name]

	if !item.Takeable {
		return item, false
	}

	delete(common.Inventory, name)

	return item, true
}

func (common Common) Info() {
	fmt.Printf("%v\n%v\n", common.Name, common.Description)
}

func (common Common) ListInventory() {
	if len(common.Inventory) == 0 {
		fmt.Println("{}")
	} else {
		for _, v := range common.Inventory {
			fmt.Printf("%s\n", v.Name)
		}
		fmt.Println("\n")
	}
}
