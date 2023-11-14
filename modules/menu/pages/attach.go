package pages

import (
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
)

type GetStdout struct {
	menu.Base
	Container string
	Follow    bool
}

func (p *GetStdout) GetHeadline() string {
	return "GetStdout"
}

func (p *GetStdout) Run() (menu.PageInterface, int, error) {
	docker.GetStdoutFromContainer(p.Container, p.Follow)

	console.WaitForReturn()

	return nil, 0, nil
}
