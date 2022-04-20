package item

type Item struct {
	Name     string
	Openable bool
	Contains Items
}

func New(name string, openable bool) *Item {
	item := Item{name, openable, make(Items)}

	return &item
}

func SetContains(items Items, names ...string) {
	for i := 1; i < len(names); i++ {
		items[names[0]].Contains[names[i]] = items[names[i]]
	}
}

type Items map[string]*Item
