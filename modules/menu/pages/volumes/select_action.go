package volumes

import (
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
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
	menuEntries := menu.MenuEntryList{}

	_, err := docker.GetVolume(p.Volume)
	if err != nil {
		return p.HandleError(err, false)
	}

	menuEntries = append(menuEntries, &menu.MenuEntry{Label: "Show Info", Page: &Inspect{Volume: p.Volume}})

	result := p.RunMenu(menuEntries, menu.MenuEntryList{})
	return p.EvaluateResult(result)
}
