package base

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/common/commands"
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type Page struct {
	commands.TaskRunner
	Args     []string
	DontPush bool
}

func (p *Page) RegisterHook() {}

func (p *Page) HasHook() bool {
	return false
}

func (p *Page) DontPushToStack() bool {
	return p.DontPush
}

func (p *Page) SetArguments(args []string) {
	p.Args = args
}

func (p *Page) HandleError(err error, confirm bool) (interfaces.PageInterface, int, error) {
	switch err.(type) {
	case console.AbortError:
		fmt.Println("Aborted")
		fmt.Println(err.Error())
		fmt.Println()

		if confirm {
			console.WaitForReturn()
		}

		return nil, 0, nil
	default:
		return nil, 1, err
	}
}

func (p *Page) EvaluateResult(result *models.Entry) (interfaces.PageInterface, int, error) {
	if nil == result {
		return nil, 0, nil
	}

	return result.Page, 0, nil
}
