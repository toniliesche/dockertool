package networks

import (
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return "Network Management"
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	networks, err := docker.FetchNetworks()
	if err != nil {
		return p.HandleError(err, true)
	}

	menuEntries := menu.MenuEntryList{}

	for _, network := range networks {
		menuEntries = append(
			menuEntries,
			&menu.MenuEntry{Label: network.Name, Page: &SelectAction{Network: network.Name}},
		)
	}

	result := p.RunMenu(menuEntries, menu.MenuEntryList{})
	return p.EvaluateResult(result)
}
