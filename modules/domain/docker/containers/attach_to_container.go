package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/application"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
)

type AttachToContainer struct {
	application.BaseCommand
	container string
	follow    bool
}

func (c *AttachToContainer) Validate() bool {
	if "" == c.container {
		c.Err = fmt.Errorf("container name must be set")

		return false
	}

	return true
}

func (c *AttachToContainer) Run() error {
	return docker.GetStdoutFromContainer(c.container, c.follow)
}

func (c *AttachToContainer) GetResult() interface{} {
	return nil
}

func CreateAttachToContainerCommand(container string, follow bool) (*AttachToContainer, error) {
	command := &AttachToContainer{container: container, follow: follow}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
