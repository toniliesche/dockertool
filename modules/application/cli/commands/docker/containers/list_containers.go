package containers

import (
	"github.com/toniliesche/dockertool/modules/application/cli/commands"
	"github.com/toniliesche/dockertool/modules/domain/tasks/docker/containers"
	"github.com/urfave/cli/v2"
)

type ListContainers struct {
	commands.Base
}

func (c *ListContainers) Run(context *cli.Context) error {
	_, err := c.CreateAndRunTask(
		containers.CreateListContainersCommand(
			context.IsSet("ro"),
			context.IsSet("so"),
			context.String("n"),
			context.String("r"),
		),
	)

	return err
}

func DefineListContainers() *cli.Command {
	cmd := &ListContainers{}

	return &cli.Command{
		Name:     "container-list",
		Category: "docker",
		Aliases:  []string{"cl"},
		Usage:    "list existing docker containers",
		Action:   cmd.Run,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "running-only", Aliases: []string{"ro"}, Usage: "mutually exclusive with stopped-only"},
			&cli.BoolFlag{Name: "stopped-only", Aliases: []string{"so"}, Usage: "mutually exclusive with running-only"},
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Usage: "filter for container name"},
			&cli.StringFlag{Name: "repo", Aliases: []string{"r"}, Usage: "filter for container repo"},
		},
	}
}
