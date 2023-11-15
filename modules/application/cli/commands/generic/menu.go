package generic

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/cli/commands/base"
	"github.com/toniliesche/dockertool/modules/application/common/library"
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/pages"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/urfave/cli/v2"
	"strings"
)

type Menu struct {
	base.Command
	Headlines library.Stack[string]
}

func (c *Menu) Run(context *cli.Context) error {
	pageObj := &pages.Index{}

	_, err := c.run(pageObj)
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

func (c *Menu) run(pageObj interfaces.PageInterface) (int, error) {
	pageStack := library.Stack[interfaces.PageInterface]{}
	c.Headlines = library.Stack[string]{}

	var returnCode int
	var err error
	var prevObj interfaces.PageInterface

	for {
		fmt.Print("\033[H\033[2J")
		headline := c.Headlines.Items()
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
			c.Headlines.Pop()

			if nil == pageObj {
				return returnCode, nil
			}
		} else if false == prevObj.DontPushToStack() {
			pageStack.Push(prevObj)
			c.Headlines.Push(prevObj.GetHeadline())
			if prevObj.HasHook() {
				prevObj.RegisterHook()
			}
		}
	}
}
