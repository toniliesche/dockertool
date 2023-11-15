package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/tasks/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
)

type AttachToContainerTask struct {
	base.Task
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

func CreateAttachToContainerCommand(container string, follow bool) (*AttachToContainerTask, error) {
	command := &AttachToContainerTask{container: container, follow: follow}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
