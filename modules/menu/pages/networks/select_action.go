package networks

import (
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
)

type SelectAction struct {
	menu.Base
	menu.Menu
	Network string
}

func (p *SelectAction) GetHeadline() string {
	return p.Network
}

func (p *SelectAction) Run() (menu.PageInterface, int, error) {
	menuEntries := menu.MenuEntryList{}

	_, err := docker.GetNetwork(p.Network)
	if err != nil {
		return p.HandleError(err, false)
	}

	menuEntries = append(menuEntries, &menu.MenuEntry{Label: "Show Info", Page: &Inspect{Network: p.Network}})

	result := p.RunMenu(menuEntries, menu.MenuEntryList{})
	return p.EvaluateResult(result)
}
