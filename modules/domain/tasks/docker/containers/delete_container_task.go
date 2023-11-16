package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/tasks/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
)

type DeleteContainerTask struct {
	base.Task
	container string
}

func (t *DeleteContainerTask) Validate() bool {
	if "" == t.container {
		t.Err = fmt.Errorf("container name must be set")

		return false
	}

	return true
}

func (t *DeleteContainerTask) Run() error {
	return containers.DeleteContainer(t.container)
}

func CreateDeleteContainerCommand(container string) (*DeleteContainerTask, error) {
	command := &DeleteContainerTask{container: container}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
