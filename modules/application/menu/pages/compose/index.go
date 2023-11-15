package compose

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
)

type SelectComposition struct {
	base.MenuPage
}

func (p *SelectComposition) GetHeadline() string {
	return "Compose Management"
}

func (p *SelectComposition) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectComposition) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	return menuEntries, specialEntries, nil
}
