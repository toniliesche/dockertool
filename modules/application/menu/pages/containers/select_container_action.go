package containers

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
)

type SelectContainerAction struct {
	base.MenuPage
	Container string
}

func (p *SelectContainerAction) GetHeadline() string {
	return p.Container
}

func (p *SelectContainerAction) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectContainerAction) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{
		{Label: "Show Info", Page: &InspectContainer{Container: p.Container}},
	}
	specialEntries := models.EntryList{}

	container, err := containers.FetchContainer(p.Container)
	if nil != err {
		return nil, nil, err
	}

	if container.IsRunning() {
		menuEntries = append(
			menuEntries,
			&models.Entry{Label: "Get Shell", Page: &GetShell{Container: p.Container}},
			&models.Entry{Label: "Attach to Stdout", Page: &GetStdout{Container: p.Container, Follow: true}},
			&models.Entry{Label: "Restart", Page: &RestartContainer{Container: p.Container}},
			&models.Entry{Label: "Stop", Page: &StopContainer{Container: p.Container}},
		)
	} else {
		menuEntries = append(
			menuEntries,
			&models.Entry{Label: "Show Stdout", Page: &GetStdout{Container: p.Container, Follow: false}},
			&models.Entry{Label: "Start", Page: &StartContainer{Container: p.Container}},
			&models.Entry{Label: "Delete", Page: &DeleteContainer{Container: p.Container}},
		)
	}

	return menuEntries, specialEntries, nil
}
