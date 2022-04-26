package game

import "fmt"

type Command struct {
	Name     string
	Help     string
	LongHelp string
	Action   func(args ...string)
}

func NewCommand(name, help, longhelp string, action func(args ...string)) *Command {
	command := Command{name, help, longhelp, action}
	return &command
}

func (command Command) ShortDescription() {
	fmt.Printf("\n%v\n%v\n", command.Name, command.Help)
}

func (command Command) LongDescription() {
	fmt.Printf("\n%v\n%v\n%v\n", command.Name, command.Help, command.LongHelp)
}

type CommandList map[string]*Command
