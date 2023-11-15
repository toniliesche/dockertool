package networks

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/networks"
)

type SelectNetworkAction struct {
	base.MenuPage
	Network string
}

func (p *SelectNetworkAction) GetHeadline() string {
	return p.Network
}

func (p *SelectNetworkAction) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectNetworkAction) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	_, err := networks.FetchNetwork(p.Network)
	if nil != err {
		return nil, nil, err
	}

	menuEntries = append(menuEntries, &models.Entry{Label: "Show Info", Page: &InspectNetwork{Network: p.Network}})

	return menuEntries, specialEntries, nil
}
