package networks

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
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

	menuEntries := menu.EntryList{}

	for _, network := range networks {
		menuEntries = append(
			menuEntries,
			&menu.Entry{Label: network.Name, Page: &SelectAction{Network: network.Name}},
		)
	}

	result := p.RunMenu(menuEntries, menu.EntryList{})
	return p.EvaluateResult(result)
}
