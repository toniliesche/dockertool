package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/tasks/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
)

type StartContainerTask struct {
	base.Task
	container string
}

func (t *StartContainerTask) Validate() bool {
	if "" == t.container {
		t.Err = fmt.Errorf("container name must be set")

		return false
	}

	return true
}

func (t *StartContainerTask) Run() error {
	return containers.StartContainer(t.container)
}

func CreateStartContainerCommand(container string) (*StartContainerTask, error) {
	command := &StartContainerTask{container: container}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
