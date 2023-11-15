package volumes

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/volumes"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return "Volume Management"
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	volumeList, err := volumes.FetchVolumeList()
	if err != nil {
		return p.HandleError(err, true)
	}

	menuEntries := menu.EntryList{}

	for _, volume := range volumeList {
		menuEntries = append(
			menuEntries,
			&menu.Entry{Label: volume.Name, Page: &SelectAction{Volume: volume.Name}},
		)
	}

	result := p.RunMenu(menuEntries, menu.EntryList{})
	return p.EvaluateResult(result)
}
