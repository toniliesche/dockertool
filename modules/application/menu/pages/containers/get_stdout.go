package containers

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/domain/tasks/docker/containers"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type GetStdout struct {
	base.Page
	Container string
	Follow    bool
}

func (p *GetStdout) GetHeadline() string {
	return "Get Stdout"
}

func (p *GetStdout) Run() (interfaces.PageInterface, int, error) {
	p.CreateAndRunTask(containers.CreateAttachToContainerCommand(p.Container, p.Follow))

	console.WaitForReturn()

	return nil, 0, nil
}
