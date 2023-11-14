package images

import (
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
)

type SelectAction struct {
	menu.Base
	menu.Menu
	ID         string
	Repository string
	Tag        string
}

func (p *SelectAction) GetHeadline() string {
	return p.Repository + ":" + p.Tag
}

func (p *SelectAction) Run() (menu.PageInterface, int, error) {
	menuEntries := menu.MenuEntryList{}

	_, err := docker.GetImage(p.ID)
	if err != nil {
		return p.HandleError(err, false)
	}

	menuEntries = append(menuEntries, &menu.MenuEntry{Label: "Show Info", Page: &Inspect{Image: p.ID}})

	result := p.RunMenu(menuEntries, menu.MenuEntryList{})
	return p.EvaluateResult(result)
}
