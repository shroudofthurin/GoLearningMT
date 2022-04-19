package item

type Item struct {
	Name     string
	Openable bool
}

type Items map[string]*Item
