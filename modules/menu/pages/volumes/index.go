package volumes

import (
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return "Volume Management"
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	volumes, err := docker.FetchVolumes()
	if err != nil {
		return p.HandleError(err, true)
	}

	menuEntries := menu.MenuEntryList{}

	for _, volume := range volumes {
		menuEntries = append(
			menuEntries,
			&menu.MenuEntry{Label: volume.Name, Page: &SelectAction{Volume: volume.Name}},
		)
	}

	result := p.RunMenu(menuEntries, menu.MenuEntryList{})
	return p.EvaluateResult(result)
}
