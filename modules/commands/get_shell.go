package commands

import (
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/urfave/cli/v2"
)

type GetShell struct {
}

func (g *GetShell) Run(context *cli.Context) error {
	return docker.GetShell(context.Args().Get(0))
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
