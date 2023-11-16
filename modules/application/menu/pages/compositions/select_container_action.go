package compositions

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/containers"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/compose"
	containers2 "github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
	"sort"
)

type SelectContainerAction struct {
	base.MenuPage
	Composition string
}

func (p *SelectContainerAction) GetHeadline() string {
	return "Containers"
}

func (p *SelectContainerAction) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectContainerAction) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	composition, err := compose.FetchComposition(p.Composition)
	if err != nil {
		return nil, nil, err
	}

	containerList, err := containers2.FetchContainerListByComposition(composition.Name, composition.ConfigFiles)
	if err != nil {
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
			&models.Entry{Label: fmt.Sprintf("%s (Running: %s)", container.Name, console.BoolToYesNoColored(container.IsRunning())), Page: &containers.SelectContainerAction{Container: container.Name}, Divider: index == dividingIndex},
		)
	}

	return menuEntries, specialEntries, nil
}
