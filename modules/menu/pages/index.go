package pages

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/menu"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return fmt.Sprintf("Docker Management Tool (%s%s%s)", console.InfoColor, "test", console.HeadlineColor)
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	menuEntries := menu.MenuEntryList{
		&menu.MenuEntry{Label: "Container Management", Page: &Containers{}},
	}

	result := p.RunMenu(menuEntries, menu.MenuEntryList{{Label: "Show Version Information", Page: &Version{}, Shortcut: "v"}})
	return p.EvaluateResult(result)
}
