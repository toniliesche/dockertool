package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/shared"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
	"sort"
)

type ListContainersTask struct {
	shared.BaseTask
	options *containers.FilterOptions
}

func (t *ListContainersTask) Validate() bool {
	return true
}

func (t *ListContainersTask) Run() error {
	containerList, err := containers.FetchContainerList(t.options)
	if err != nil {
		return err
	}
	mapping := map[string]int{}
	keys := make([]string, 0, len(containerList))
	keysStopped := make([]string, 0, len(containerList))
	for key, container := range containerList {
		keys = append(keys, container.Name)
		mapping[container.Name] = key
	}

	sort.Strings(keys)
	keys = append(keys, keysStopped...)

	for _, key := range keys {
		container := containerList[mapping[key]]
		fmt.Printf("container : %s (running : %s)\n", container.Name, container.IsRunningString())
	}

	return nil
}

func (t *ListContainersTask) GetResult() interface{} {
	return nil
}

func CreateListContainersCommand(runningOnly bool, stoppedOnly bool, nameFilter string, repoFilter string) (*ListContainersTask, error) {
	options := &containers.FilterOptions{
		StateFilter: runningOnly || stoppedOnly,
		RunningOnly: runningOnly,
		StoppedOnly: stoppedOnly,
		NameFilter:  nameFilter,
		RepoFilter:  repoFilter,
	}

	return &ListContainersTask{options: options}, nil
}
