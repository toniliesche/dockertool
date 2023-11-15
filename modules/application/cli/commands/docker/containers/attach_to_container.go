package containers

import (
	"github.com/toniliesche/dockertool/modules/application/cli/commands"
	"github.com/toniliesche/dockertool/modules/domain/docker/containers"
	"github.com/urfave/cli/v2"
)

type AttachToContainer struct {
	commands.Base
}

func (c *AttachToContainer) Run(context *cli.Context) error {
	_, err := c.CreateRunCommand(containers.CreateAttachToContainerCommand(context.Args().Get(0), true))

	return err
}

func DefineAttachToContainer() *cli.Command {
	cmd := &AttachToContainer{}

	return &cli.Command{
		Name:      "attach",
		Aliases:   []string{"a"},
		Usage:     "attach to existing container",
		ArgsUsage: "[container]",
		Action:    cmd.Run,
	}
}
