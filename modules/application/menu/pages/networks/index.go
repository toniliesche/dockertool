package networks

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/networks"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return "Network Management"
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	networkList, err := networks.FetchNetworkList()
	if err != nil {
		return p.HandleError(err, true)
	}

	menuEntries := menu.EntryList{}

	for _, network := range networkList {
		menuEntries = append(
			menuEntries,
			&menu.Entry{Label: network.Name, Page: &SelectAction{Network: network.Name}},
		)
	}

	result := p.RunMenu(menuEntries, menu.EntryList{})
	return p.EvaluateResult(result)
}
