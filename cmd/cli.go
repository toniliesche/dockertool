package main

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/commands"
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/state"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

func main() {
	state.CreateState()
	commandFactory := commands.GetCommandFactory()

	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "dockertool",
		Compiled:             time.Now(),
		Authors: []*cli.Author{
			{
				Name:  state.AppState.AuthorName,
				Email: state.AppState.AuthorMail,
			},
		},
		Copyright: fmt.Sprintf("(c) %s %s", state.AppState.CopyrightYear, state.AppState.Copyright),
		Usage:     "manage docker containers from local terminal",
		Commands:  commandFactory.GetCommands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		console.PrintError(err.Error())
		os.Exit(1)
	}
}
