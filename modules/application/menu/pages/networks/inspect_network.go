package networks

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/domain/tasks/docker/networks"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type InspectNetwork struct {
	base.Page
	Network string
}

func (p *InspectNetwork) GetHeadline() string {
	return "Show Info"
}

func (p *InspectNetwork) Run() (interfaces.PageInterface, int, error) {
	_, err := p.CreateAndRunTask(networks.CreateInspectNetworkCommand(p.Network))
	if nil != err {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
