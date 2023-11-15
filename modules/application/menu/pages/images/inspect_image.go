package images

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/domain/tasks/docker/images"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type InspectImage struct {
	base.Page
	Image string
}

func (p *InspectImage) GetHeadline() string {
	return "Show Info"
}

func (p *InspectImage) Run() (interfaces.PageInterface, int, error) {
	_, err := p.CreateAndRunTask(images.CreateInspectImageCommand(p.Image))
	if nil != err {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
