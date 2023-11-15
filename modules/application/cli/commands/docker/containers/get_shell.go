package containers

import (
	"github.com/toniliesche/dockertool/modules/application/cli/commands"
	"github.com/toniliesche/dockertool/modules/domain/docker/containers"
	"github.com/urfave/cli/v2"
)

type GetShell struct {
	commands.Base
}

func (c *GetShell) Run(context *cli.Context) error {
	_, err := c.CreateRunCommand(containers.CreateGetContainerShellCommand(context.Args().Get(0)))

	return err
}

func DefineGetShell() *cli.Command {
	cmd := &GetShell{}

	return &cli.Command{
		Name:      "shell",
		Aliases:   []string{"s"},
		Usage:     "get shell inside specified container",
		ArgsUsage: "[container]",
		Action:    cmd.Run,
	}
}