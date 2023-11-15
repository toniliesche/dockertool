package networks

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/domain/docker/networks"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type Inspect struct {
	menu.Base
	Network string
}

func (p *Inspect) GetHeadline() string {
	return "Show Info"
}

func (p *Inspect) Run() (menu.PageInterface, int, error) {
	_, err := p.CreateRunCommand(networks.CreateInspectNetworkCommand(p.Network))
	if err != nil {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
