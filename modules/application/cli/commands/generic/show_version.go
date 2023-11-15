package generic

import (
	"github.com/toniliesche/dockertool/modules/application/cli/commands/base"
	"github.com/toniliesche/dockertool/modules/domain/tasks/generic"
	"github.com/urfave/cli/v2"
)

type ShowVersion struct {
	base.Command
}

func (c *ShowVersion) Run(context *cli.Context) error {
	_, err := c.CreateAndRunTask(generic.CreateShowVersionInformationCommand())

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
