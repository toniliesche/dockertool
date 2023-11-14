package commands

import (
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/urfave/cli/v2"
)

type AttachToContainer struct {
}

func (a *AttachToContainer) Run(context *cli.Context) error {
	return docker.GetStdoutFromContainer(context.Args().Get(0), true)
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
