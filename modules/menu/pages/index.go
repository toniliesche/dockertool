package pages

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/menu"
	"github.com/toniliesche/dockertool/modules/menu/pages/compose"
	"github.com/toniliesche/dockertool/modules/menu/pages/containers"
	"github.com/toniliesche/dockertool/modules/menu/pages/generic"
	"github.com/toniliesche/dockertool/modules/menu/pages/images"
	"github.com/toniliesche/dockertool/modules/menu/pages/networks"
	"github.com/toniliesche/dockertool/modules/menu/pages/volumes"
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
		&menu.MenuEntry{Label: "Compose Management", Page: &compose.Index{}},
		&menu.MenuEntry{Label: "Container Management", Page: &containers.Index{}},
		&menu.MenuEntry{Label: "Image Management", Page: &images.Index{}},
		&menu.MenuEntry{Label: "Network Management", Page: &networks.Index{}},
		&menu.MenuEntry{Label: "Volume Management", Page: &volumes.Index{}},
	}

	result := p.RunMenu(menuEntries, menu.MenuEntryList{{Label: "Show Version Information", Page: &generic.Version{}, Shortcut: "v"}})
	return p.EvaluateResult(result)
}
