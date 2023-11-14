package containers

import (
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
)

type SelectAction struct {
	menu.Base
	menu.Menu
	Container string
}

func (p *SelectAction) GetHeadline() string {
	return p.Container
}

func (p *SelectAction) Run() (menu.PageInterface, int, error) {
	menuEntries := menu.MenuEntryList{}

	container, err := docker.GetContainer(p.Container)
	if err != nil {
		return p.HandleError(err, false)
	}

	if container.IsRunning() {
		menuEntries = append(
			menuEntries,
			&menu.MenuEntry{Label: "Get Shell", Page: &GetShell{Container: p.Container}},
			&menu.MenuEntry{Label: "Attach to Stdout", Page: &GetStdout{Container: p.Container, Follow: true}},
		)
	} else {
		menuEntries = append(
			menuEntries,
			&menu.MenuEntry{Label: "Show Stdout", Page: &GetStdout{Container: p.Container, Follow: false}},
		)
	}
	menuEntries = append(menuEntries, &menu.MenuEntry{Label: "Show Info", Page: &Inspect{Container: p.Container}})

	result := p.RunMenu(menuEntries, menu.MenuEntryList{})
	return p.EvaluateResult(result)
}
