package commands

import (
	"github.com/toniliesche/dockertool/modules/shared"
	"github.com/urfave/cli/v2"
)

type ShowVersion struct {
}

func (l *ShowVersion) Run(context *cli.Context) error {
	shared.ShowVersion()

	return nil
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
