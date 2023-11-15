package networks

import (
	menu2 "github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/networks"
)

type SelectAction struct {
	menu2.Base
	menu2.Menu
	Network string
}

func (p *SelectAction) GetHeadline() string {
	return p.Network
}

func (p *SelectAction) Run() (menu2.PageInterface, int, error) {
	menuEntries := menu2.EntryList{}

	_, err := networks.FetchNetwork(p.Network)
	if err != nil {
		return p.HandleError(err, false)
	}

	menuEntries = append(menuEntries, &menu2.Entry{Label: "Show Info", Page: &Inspect{Network: p.Network}})

	result := p.RunMenu(menuEntries, menu2.EntryList{})
	return p.EvaluateResult(result)
}
