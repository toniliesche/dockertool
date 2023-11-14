package pages

import (
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
)

type GetShell struct {
	menu.Base
	Container string
}

func (p *GetShell) GetHeadline() string {
	return "Get Shell"
}

func (p *GetShell) Run() (menu.PageInterface, int, error) {
	err := docker.GetShell(p.Container)
	if err != nil {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
