package volumes

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/domain/tasks/docker/volumes"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type Inspect struct {
	menu.Base
	Volume string
}

func (p *Inspect) GetHeadline() string {
	return "Show Info"
}

func (p *Inspect) Run() (menu.PageInterface, int, error) {
	_, err := p.CreateAndRunTask(volumes.CreateInspectVolumeCommand(p.Volume))
	if err != nil {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
