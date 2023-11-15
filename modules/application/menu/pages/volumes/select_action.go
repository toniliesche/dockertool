package volumes

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
)

type SelectAction struct {
	menu.Base
	menu.Menu
	Volume string
}

func (p *SelectAction) GetHeadline() string {
	return p.Volume
}

func (p *SelectAction) Run() (menu.PageInterface, int, error) {
	menuEntries := menu.EntryList{}

	_, err := docker.GetVolume(p.Volume)
	if err != nil {
		return p.HandleError(err, false)
	}

	menuEntries = append(menuEntries, &menu.Entry{Label: "Show Info", Page: &Inspect{Volume: p.Volume}})

	result := p.RunMenu(menuEntries, menu.EntryList{})
	return p.EvaluateResult(result)
}
