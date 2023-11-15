package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/tasks/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
)

type GetContainerShellTask struct {
	base.Task
	container string
}

func (t *GetContainerShellTask) Validate() bool {
	if "" == t.container {
		t.Err = fmt.Errorf("container name must be set")

		return false
	}

	return true
}

func (t *GetContainerShellTask) Run() error {
	return containers.GetShell(t.container)
}

func CreateGetContainerShellCommand(container string) (*GetContainerShellTask, error) {
	command := &GetContainerShellTask{container: container}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
