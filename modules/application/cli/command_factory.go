package cli

import (
	"github.com/toniliesche/dockertool/modules/application/cli/commands/docker/containers"
	"github.com/toniliesche/dockertool/modules/application/cli/commands/shared"
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
			shared.DefineMenu(),
			containers.DefineGetShell(),
			shared.DefineShowVersion(),
		}

		factory = &CommandFactory{commands: commands}
	}

	return factory
}
