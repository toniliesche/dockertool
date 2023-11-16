package compositions

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/compose"
)

type SelectComposition struct {
	base.MenuPage
}

func (p *SelectComposition) GetHeadline() string {
	return "Compositions"
}

func (p *SelectComposition) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectComposition) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	compositions, err := compose.FetchCompositionList()
	if nil != err {
		return nil, nil, err
	}

	for _, composition := range compositions {
		menuEntries = append(menuEntries, &models.Entry{Label: fmt.Sprintf("%s (Running: %s)", composition.Name, p.evaluateRunStatus(composition)), Page: &SelectCompositionAction{Composition: composition.Name}})
	}

	return menuEntries, specialEntries, nil
}

func (p *SelectComposition) evaluateRunStatus(composition *compose.Composition) string {
	if (composition.Running > 0) && (composition.Exited > 0) {
		return fmt.Sprintf("%spartly%s", console.HeadlineColor, console.RegularColor)
	}

	if (composition.Running) > 0 {
		return fmt.Sprintf("%syes%s", console.OKColor, console.RegularColor)
	}

	return fmt.Sprintf("%sno%s", console.ErrorColor, console.RegularColor)
}
