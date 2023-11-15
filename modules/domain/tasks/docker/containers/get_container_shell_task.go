package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/shared"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
)

type GetContainerShellTask struct {
	shared.BaseTask
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

func (t *GetContainerShellTask) GetResult() interface{} {
	return nil
}

func CreateGetContainerShellCommand(container string) (*GetContainerShellTask, error) {
	command := &GetContainerShellTask{container: container}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
