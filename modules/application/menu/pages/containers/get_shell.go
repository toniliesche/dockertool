package containers

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/domain/tasks/docker/containers"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type GetShell struct {
	menu.Base
	Container string
}

func (p *GetShell) GetHeadline() string {
	return "Get Shell"
}

func (p *GetShell) Run() (menu.PageInterface, int, error) {
	_, err := p.CreateAndRunTask(containers.CreateGetContainerShellCommand(p.Container))
	if err != nil {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
