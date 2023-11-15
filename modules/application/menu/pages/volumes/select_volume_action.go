package volumes

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/volumes"
)

type SelectVolumeAction struct {
	base.MenuPage
	Volume string
}

func (p *SelectVolumeAction) GetHeadline() string {
	return p.Volume
}

func (p *SelectVolumeAction) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectVolumeAction) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	_, err := volumes.FetchVolume(p.Volume)
	if nil != err {
		return nil, nil, err
	}

	menuEntries = append(menuEntries, &models.Entry{Label: "Show Info", Page: &InspectVolume{Volume: p.Volume}})

	return menuEntries, specialEntries, nil
}
