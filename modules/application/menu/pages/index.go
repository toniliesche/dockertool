package pages

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/compose"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/containers"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/generic"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/images"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/networks"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/volumes"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
)

type Index struct {
	base.MenuPage
}

func (p *Index) GetHeadline() string {
	return fmt.Sprintf("Docker Management Tool (%s%s%s)", console.InfoColor, "test", console.HeadlineColor)
}

func (p *Index) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *Index) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{
		{Label: "Compose Management", Page: &compose.SelectComposition{}},
		{Label: "Container Management", Page: &containers.SelectContainer{}},
		{Label: "Image Management", Page: &images.SelectImage{}},
		{Label: "Network Management", Page: &networks.SelectNetwork{}},
		{Label: "Volume Management", Page: &volumes.SelectVolume{}},
	}

	specialEntries := models.EntryList{
		{Label: "Show Version Information", Page: &generic.Version{}, Shortcut: "v"},
	}

	return menuEntries, specialEntries, nil
}
