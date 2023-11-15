package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
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
	containerList, err := containers.FetchContainerList()
	if err != nil {
		return p.HandleError(err, true)
	}

	menuEntries := menu.EntryList{}

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
			&menu.Entry{Label: fmt.Sprintf("%s (Running: %s)", container.Name, console.BoolToYesNoColored(container.IsRunning())), Page: &SelectAction{Container: container.Name}, Divider: index == dividingIndex},
		)
	}

	result := p.RunMenu(menuEntries, menu.EntryList{})
	return p.EvaluateResult(result)
}
