package containers

import (
	"github.com/toniliesche/dockertool/modules/application/cli/commands/base"
	"github.com/toniliesche/dockertool/modules/domain/tasks/docker/containers"
	"github.com/urfave/cli/v2"
)

type AttachToContainer struct {
	base.Command
}

func (c *AttachToContainer) Run(context *cli.Context) error {
	_, err := c.CreateAndRunTask(containers.CreateAttachToContainerCommand(context.Args().Get(0), true))

	return err
}

func DefineAttachToContainer() *cli.Command {
	cmd := &AttachToContainer{}

	return &cli.Command{
		Name:      "container-attach",
		Category:  "docker",
		Aliases:   []string{"ca"},
		Usage:     "attach to existing container",
		ArgsUsage: "[container]",
		Action:    cmd.Run,
	}
}
