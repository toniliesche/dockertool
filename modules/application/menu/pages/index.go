package pages

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/compose"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/containers"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/generic"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/images"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/networks"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/volumes"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return fmt.Sprintf("Docker Management Tool (%s%s%s)", console.InfoColor, "test", console.HeadlineColor)
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	menuEntries := menu.EntryList{
		&menu.Entry{Label: "Compose Management", Page: &compose.Index{}},
		&menu.Entry{Label: "Container Management", Page: &containers.Index{}},
		&menu.Entry{Label: "Image Management", Page: &images.Index{}},
		&menu.Entry{Label: "Network Management", Page: &networks.Index{}},
		&menu.Entry{Label: "Volume Management", Page: &volumes.Index{}},
	}

	result := p.RunMenu(menuEntries, menu.EntryList{{Label: "Show Version Information", Page: &generic.Version{}, Shortcut: "v"}})
	return p.EvaluateResult(result)
}
