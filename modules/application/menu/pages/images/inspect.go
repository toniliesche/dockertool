package images

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/domain/docker/images"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type Inspect struct {
	menu.Base
	Image string
}

func (p *Inspect) GetHeadline() string {
	return "Show Info"
}

func (p *Inspect) Run() (menu.PageInterface, int, error) {
	_, err := p.CreateRunCommand(images.CreateInspectImageCommand(p.Image))
	if err != nil {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
