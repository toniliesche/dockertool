package commands

import (
	"github.com/urfave/cli/v2"
)

var factory *CommandFactory

type CommandFactory struct {
	commands []*cli.Command
}

func (c *CommandFactory) GetCommands() []*cli.Command {
	return c.commands
}

func GetCommandFactory() *CommandFactory {
	if nil == factory {
		commands := []*cli.Command{
			DefineAttachToContainer(),
			DefineListContainers(),
			DefineMenu(),
			DefineGetShell(),
			DefineShowVersion(),
		}

		factory = &CommandFactory{commands: commands}
	}

	return factory
}
