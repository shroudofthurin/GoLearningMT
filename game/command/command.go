package command

type Command struct {
	Name     string
	Help     string
	LongHelp string
	Action   func(args string)
}

func New(name, help, longhelp string, action func(args string)) *Command {
	command := Command{name, help, longhelp, action}
	return &command
}

type CommandList map[string]*Command
