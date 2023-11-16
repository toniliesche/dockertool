package containers

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/domain/tasks/docker/containers"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type StartContainer struct {
	base.Page
	Container string
}

func (p *StartContainer) GetHeadline() string {
	return "Start Composition"
}

func (p *StartContainer) Run() (interfaces.PageInterface, int, error) {
	_, err := p.CreateAndRunTask(containers.CreateStartContainerCommand(p.Container))
	if nil != err {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
