package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/tasks/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
)

type RestartContainerTask struct {
	base.Task
	container string
}

func (t *RestartContainerTask) Validate() bool {
	if "" == t.container {
		t.Err = fmt.Errorf("container name must be set")

		return false
	}

	return true
}

func (t *RestartContainerTask) Run() error {
	return containers.RestartContainer(t.container)
}

func CreateRestartContainerCommand(container string) (*RestartContainerTask, error) {
	command := &RestartContainerTask{container: container}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
