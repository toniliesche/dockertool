package compose

import (
	"github.com/toniliesche/dockertool/modules/menu"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return "Compose Management"
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	menuEntries := menu.MenuEntryList{}

	result := p.RunMenu(menuEntries, menu.MenuEntryList{})
	return p.EvaluateResult(result)
}
