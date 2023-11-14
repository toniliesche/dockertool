package commands

import "github.com/urfave/cli/v2"

type Command interface {
	Name() string
	Run(context *cli.Context) error
	RunHelp() error
}

type CommandDefinition struct {
	Name        string
	Description string
	Construct   func(args []string) (Command, error)
}
