package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"sort"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return "Container Management"
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	containers, err := docker.FetchContainers()
	if err != nil {
		return p.HandleError(err, true)
	}

	menuEntries := menu.EntryList{}

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
			&menu.Entry{Label: fmt.Sprintf("%s (Running: %s)", container.Name, console.BoolToYesNoColored(container.IsRunning())), Page: &SelectAction{Container: container.Name}, Divider: index == dividingIndex},
		)
	}

	result := p.RunMenu(menuEntries, menu.EntryList{})
	return p.EvaluateResult(result)
}
