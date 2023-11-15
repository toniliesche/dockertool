package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/shared"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
)

type AttachToContainerTask struct {
	shared.BaseTask
	container string
	follow    bool
}

func (t *AttachToContainerTask) Validate() bool {
	if "" == t.container {
		t.Err = fmt.Errorf("container name must be set")

		return false
	}

	return true
}

func (t *AttachToContainerTask) Run() error {
	return containers.GetStdoutFromContainer(t.container, t.follow)
}

func (t *AttachToContainerTask) GetResult() interface{} {
	return nil
}

func CreateAttachToContainerCommand(container string, follow bool) (*AttachToContainerTask, error) {
	command := &AttachToContainerTask{container: container, follow: follow}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
