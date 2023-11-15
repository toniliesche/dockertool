package volumes

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/domain/tasks/docker/volumes"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type InspectVolume struct {
	base.Page
	Volume string
}

func (p *InspectVolume) GetHeadline() string {
	return "Show Info"
}

func (p *InspectVolume) Run() (interfaces.PageInterface, int, error) {
	_, err := p.CreateAndRunTask(volumes.CreateInspectVolumeCommand(p.Volume))
	if nil != err {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
