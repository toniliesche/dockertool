package networks

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/networks"
)

type SelectNetwork struct {
	base.MenuPage
}

func (p *SelectNetwork) GetHeadline() string {
	return "Networks"
}

func (p *SelectNetwork) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectNetwork) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	networkList, err := networks.FetchNetworkList()
	if nil != err {
		return nil, nil, err
	}

	for _, network := range networkList {
		menuEntries = append(
			menuEntries,
			&models.Entry{Label: network.Name, Page: &SelectNetworkAction{Network: network.Name}},
		)
	}

	return menuEntries, specialEntries, nil
}
