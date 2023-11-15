package generic

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/domain/shared"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type Version struct {
	menu.Base
}

func (p *Version) GetHeadline() string {
	return "Show Menu Version"
}

func (p *Version) Run() (menu.PageInterface, int, error) {
	_, err := p.CreateRunCommand(shared.CreateShowVersionInformationCommand())
	if err != nil {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
