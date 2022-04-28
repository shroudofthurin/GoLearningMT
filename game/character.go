package game

import "fmt"

type Character struct {
	Common
	Phrase  string
	Askable bool
}

func NewCharacter(name, description, phrase string, points int, askable bool) *Character {
	character := Character{
		Common{name, description, points, make(Items)},
		phrase,
		askable,
	}

	return &character
}

func (character *Character) Take(name string) (*Item, bool) {
	item := character.Inventory[name]

	delete(character.Inventory, name)

	if len(character.Inventory) == 0 {
		character.Askable = false
	}

	return item, true
}

func (character *Character) Info() {
	fmt.Printf("%v\n%v\n", character.Name, character.Description)
	if character.Askable {
		fmt.Println("\nInventory:")
		character.ListInventory()
	}
}

type Characters map[string]*Character

func SetCharacterInventory(characters Characters, items Items, names ...string) {
	for i := 1; i < len(names); i++ {
		characters[names[0]].Inventory[names[i]] = items[names[i]]
	}
}
