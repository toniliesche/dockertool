package compose

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return "Compose Management"
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	menuEntries := menu.EntryList{}

	result := p.RunMenu(menuEntries, menu.EntryList{})
	return p.EvaluateResult(result)
}
