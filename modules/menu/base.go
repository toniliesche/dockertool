package menu

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/console"
)

type Base struct {
	Args     []string
	DontPush bool
}

func (p *Base) RegisterHook() {}

func (p *Base) HasHook() bool {
	return false
}

func (p *Base) DontPushToStack() bool {
	return p.DontPush
}

func (p *Base) SetArguments(args []string) {
	p.Args = args
}

func (p *Base) HandleError(err error, confirm bool) (PageInterface, int, error) {
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

func (p *Base) EvaluateResult(result *MenuEntry) (PageInterface, int, error) {
	if nil == result {
		return nil, 0, nil
	}

	return result.Page, 0, nil
}
