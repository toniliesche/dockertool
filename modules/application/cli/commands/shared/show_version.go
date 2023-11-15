package shared

import (
	"github.com/toniliesche/dockertool/modules/application/cli/commands"
	"github.com/toniliesche/dockertool/modules/domain/shared"
	"github.com/urfave/cli/v2"
)

type ShowVersion struct {
	commands.Base
}

func (c *ShowVersion) Run(context *cli.Context) error {
	_, err := c.CreateRunCommand(shared.CreateShowVersionInformationCommand())

	return err
}

func DefineShowVersion() *cli.Command {
	cmd := &ShowVersion{}

	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show version information",
		Action:  cmd.Run,
	}
}
