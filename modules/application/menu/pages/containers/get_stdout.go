package containers

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/domain/docker/containers"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type GetStdout struct {
	menu.Base
	Container string
	Follow    bool
}

func (p *GetStdout) GetHeadline() string {
	return "Get Stdout"
}

func (p *GetStdout) Run() (menu.PageInterface, int, error) {
	p.CreateRunCommand(containers.CreateAttachToContainerCommand(p.Container, p.Follow))

	console.WaitForReturn()

	return nil, 0, nil
}
