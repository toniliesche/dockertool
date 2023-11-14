package pages

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
	"sort"
)

type Containers struct {
	menu.Base
	menu.Menu
}

func (p *Containers) GetHeadline() string {
	return "Container Management"
}

func (p *Containers) Run() (menu.PageInterface, int, error) {
	containers, err := docker.FetchContainers()
	if err != nil {
		return p.HandleError(err, true)
	}

	menuEntries := menu.MenuEntryList{}

	mapping := map[string]int{}
	keys := make([]string, 0, len(containers))
	keysStopped := make([]string, 0, len(containers))
	for key, container := range containers {
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
		container := containers[mapping[key]]
		menuEntries = append(
			menuEntries,
			&menu.MenuEntry{Label: fmt.Sprintf("%s (Running: %s)", container.Name, console.BoolToYesNoColored(container.IsRunning())), Page: &ContainerActions{Container: container.Name}, Divider: index == dividingIndex},
		)
	}

	result := p.RunMenu(menuEntries, menu.MenuEntryList{})
	return p.EvaluateResult(result)
}
