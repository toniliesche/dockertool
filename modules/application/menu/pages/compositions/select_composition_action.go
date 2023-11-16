package compositions

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/compose"
)

type SelectCompositionAction struct {
	base.MenuPage
	Composition string
}

func (p *SelectCompositionAction) GetHeadline() string {
	return p.Composition
}

func (p *SelectCompositionAction) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectCompositionAction) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{
		{Label: "Containers", Page: &SelectContainerAction{Composition: p.Composition}},
	}
	specialEntries := models.EntryList{}

	composition, err := compose.FetchComposition(p.Composition)
	if err != nil {
		return nil, nil, err
	}

	if !composition.IsFullyRunning() {
		menuEntries = append(menuEntries, &models.Entry{Label: "Start"})
	} else {
		menuEntries = append(menuEntries, &models.Entry{Label: "Restart"})
	}

	if !composition.IsFullyStopped() {
		menuEntries = append(menuEntries, &models.Entry{Label: "Stop"})
	}

	return menuEntries, specialEntries, nil
}
