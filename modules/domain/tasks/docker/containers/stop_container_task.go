package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/tasks/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
)

type StopContainerTask struct {
	base.Task
	container string
}

func (t *StopContainerTask) Validate() bool {
	if "" == t.container {
		t.Err = fmt.Errorf("container name must be set")

		return false
	}

	return true
}

func (t *StopContainerTask) Run() error {
	return containers.StopContainer(t.container)
}

func CreateStopContainerCommand(container string) (*StopContainerTask, error) {
	command := &StopContainerTask{container: container}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
