package commands

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/library"
	"github.com/toniliesche/dockertool/modules/menu"
	"github.com/toniliesche/dockertool/modules/menu/pages"
	"github.com/toniliesche/dockertool/modules/state"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Menu struct {
	Headlines library.Stack[string]
}

func (m *Menu) Run(context *cli.Context) error {
	pageObj := &pages.Index{}

	_, err := m.run(pageObj)
	return err
}

func DefineMenu() *cli.Command {
	cmd := &Menu{}

	return &cli.Command{
		Name:    "menu",
		Aliases: []string{"m"},
		Usage:   "menu navigation",
		Action:  cmd.Run,
	}
}

func (m *Menu) run(pageObj menu.PageInterface) (int, error) {
	pageStack := library.Stack[menu.PageInterface]{}
	m.Headlines = library.Stack[string]{}

	state.AppState.ShutdownChannel = make(chan os.Signal)
	defer close(state.AppState.ShutdownChannel)
	signal.Notify(state.AppState.ShutdownChannel, os.Interrupt, syscall.SIGTERM)
	defer signal.Reset(os.Interrupt, syscall.SIGTERM)

	go func() {
		<-state.AppState.ShutdownChannel
		m.shutdown()
		os.Exit(1)
	}()

	var returnCode int
	var err error
	var prevObj menu.PageInterface

	for {
		fmt.Print("\033[H\033[2J")
		headline := m.Headlines.Items()
		headline = append(headline, pageObj.GetHeadline())
		console.PrintHeadline(strings.Join(headline, " > "))

		prevObj = pageObj
		pageObj, returnCode, err = pageObj.Run()
		if nil != err {
			return returnCode, fmt.Errorf("unexpected error: %s", err.Error())
		}

		if nil == pageObj {
			if pageStack.Empty() {
				return returnCode, nil
			}
			pageObj = pageStack.Pop()
			m.Headlines.Pop()

			if nil == pageObj {
				return returnCode, nil
			}
		} else if false == prevObj.DontPushToStack() {
			pageStack.Push(prevObj)
			m.Headlines.Push(prevObj.GetHeadline())
			if prevObj.HasHook() {
				prevObj.RegisterHook()
			}
		}
	}
}

func (m *Menu) shutdown() {

}
