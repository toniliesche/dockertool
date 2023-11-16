package volumes

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/volumes"
)

type SelectVolume struct {
	base.MenuPage
}

func (p *SelectVolume) GetHeadline() string {
	return "Volumes"
}

func (p *SelectVolume) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectVolume) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	volumeList, err := volumes.FetchVolumeList()
	if nil != err {
		return nil, nil, err
	}

	for _, volume := range volumeList {
		menuEntries = append(
			menuEntries,
			&models.Entry{Label: volume.Name, Page: &SelectVolumeAction{Volume: volume.Name}},
		)
	}

	return menuEntries, specialEntries, nil
}
