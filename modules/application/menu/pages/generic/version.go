package generic

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/domain/tasks/generic"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type Version struct {
	base.Page
}

func (p *Version) GetHeadline() string {
	return "Show MenuPage Version"
}

func (p *Version) Run() (interfaces.PageInterface, int, error) {
	_, err := p.CreateAndRunTask(generic.CreateShowVersionInformationCommand())
	if nil != err {
		return p.HandleError(err, true)
	}

	console.WaitForReturn()

	return nil, 0, nil
}
