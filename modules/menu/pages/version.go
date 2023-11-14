package pages

import (
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/menu"
	"github.com/toniliesche/dockertool/modules/shared"
)

type Version struct {
	menu.Base
}

func (p *Version) GetHeadline() string {
	return "Show Menu Version"
}

func (p *Version) Run() (menu.PageInterface, int, error) {
	shared.ShowVersion()

	console.WaitForReturn()

	return nil, 0, nil
}
