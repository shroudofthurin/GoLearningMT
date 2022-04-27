package game

type Character struct {
	Common
	Phrase  string
	Askable bool
}

func NewCharacter(name, description, phrase string, askable bool) *Character {
	character := Character{
		Common{name, description, 0, make(Items)},
		phrase,
		askable,
	}

	return &character
}

type Characters map[string]*Character

func SetCharacterInventory(characters Characters, items Items, names ...string) {
	for i := 1; i < len(names); i++ {
		characters[names[0]].Inventory[names[i]] = items[names[i]]
	}
}
