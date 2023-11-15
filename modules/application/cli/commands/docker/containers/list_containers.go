package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/cli/commands"
	docker2 "github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"github.com/urfave/cli/v2"
)

type ListContainers struct {
	commands.Base
}

func (c *ListContainers) Run(context *cli.Context) error {
	options := &docker2.FilterOptions{
		StateFilter: context.IsSet("ro") || context.IsSet("so"),
		RunningOnly: context.IsSet("ro"),
		StoppedOnly: context.IsSet("so"),
		NameFilter:  context.String("n"),
		RepoFilter:  context.String("r"),
	}

	containers, err := docker2.FetchContainers(options)
	if err != nil {
		return err
	}

	for _, container := range containers {
		fmt.Printf("container : %s (running : %s)\n", container.Name, container.IsRunningString())
	}

	return nil
}

func DefineListContainers() *cli.Command {
	cmd := &ListContainers{}

	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list existing docker containers",
		Action:  cmd.Run,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "running-only", Aliases: []string{"ro"}, Usage: "mutually exclusive with stopped-only"},
			&cli.BoolFlag{Name: "stopped-only", Aliases: []string{"so"}, Usage: "mutually exclusive with running-only"},
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Usage: "filter for container name"},
			&cli.StringFlag{Name: "repo", Aliases: []string{"r"}, Usage: "filter for container repo"},
		},
	}
}
