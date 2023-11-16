package main

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application"
	application_cli "github.com/toniliesche/dockertool/modules/application/cli/commands/factory"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/system"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	application.CreateState()
	commandFactory := application_cli.GetCommandFactory()

	application.AppState.ShutdownChannel = make(chan os.Signal)
	defer close(application.AppState.ShutdownChannel)
	signal.Notify(application.AppState.ShutdownChannel, os.Interrupt, syscall.SIGTERM)
	defer signal.Reset(os.Interrupt, syscall.SIGTERM)

	go func() {
		<-application.AppState.ShutdownChannel
		shutdown()
		os.Exit(1)
	}()

	err := system.CheckComposePlugin()
	application.AppState.Compose = nil == err

	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "dockertool",
		Compiled:             time.Now(),
		Authors: []*cli.Author{
			{
				Name:  application.AppState.AuthorName,
				Email: application.AppState.AuthorMail,
			},
		},
		Copyright: fmt.Sprintf("(c) %s %s", application.AppState.CopyrightYear, application.AppState.Copyright),
		Usage:     "manage docker containers from local terminal",
		Commands:  commandFactory.GetCommands(),
	}

	err = app.Run(os.Args)
	if err != nil {
		console.PrintError(err.Error())
		os.Exit(1)
	}
}

func shutdown() {

}
