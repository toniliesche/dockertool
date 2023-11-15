package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
	"sort"
)

type SelectContainer struct {
	base.MenuPage
}

func (p *SelectContainer) GetHeadline() string {
	return "Container Management"
}

func (p *SelectContainer) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectContainer) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	containerList, err := containers.FetchContainerList()
	if nil != err {
		return nil, nil, err
	}

	mapping := map[string]int{}
	keys := make([]string, 0, len(containerList))
	keysStopped := make([]string, 0, len(containerList))
	for key, container := range containerList {
		if container.IsRunning() {
			keys = append(keys, container.Name)
		} else {
			keysStopped = append(keysStopped, container.Name)
		}
		mapping[container.Name] = key
	}

	sort.Strings(keys)
	sort.Strings(keysStopped)
	dividingIndex := len(keys) - 1

	keys = append(keys, keysStopped...)

	for index, key := range keys {
		container := containerList[mapping[key]]
		menuEntries = append(
			menuEntries,
			&models.Entry{Label: fmt.Sprintf("%s (Running: %s)", container.Name, console.BoolToYesNoColored(container.IsRunning())), Page: &SelectContainerAction{Container: container.Name}, Divider: index == dividingIndex},
		)
	}

	return menuEntries, specialEntries, nil
}
