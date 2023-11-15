package images

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
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
	menuEntries := menu.EntryList{}

	_, err := docker.GetImage(p.ID)
	if err != nil {
		return p.HandleError(err, false)
	}

	menuEntries = append(menuEntries, &menu.Entry{Label: "Show Info", Page: &Inspect{Image: p.ID}})

	result := p.RunMenu(menuEntries, menu.EntryList{})
	return p.EvaluateResult(result)
}
