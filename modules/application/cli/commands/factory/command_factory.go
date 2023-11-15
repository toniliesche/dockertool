package factory

import (
	"github.com/toniliesche/dockertool/modules/application/cli/commands/docker/containers"
	"github.com/toniliesche/dockertool/modules/application/cli/commands/generic"
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
			containers.DefineAttachToContainer(),
			containers.DefineListContainers(),
			generic.DefineMenu(),
			containers.DefineGetShell(),
			generic.DefineShowVersion(),
		}

		factory = &CommandFactory{commands: commands}
	}

	return factory
}
