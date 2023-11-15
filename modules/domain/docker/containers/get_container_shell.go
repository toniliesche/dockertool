package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/application"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
)

type GetContainerShell struct {
	application.BaseCommand
	container string
}

func (c *GetContainerShell) Validate() bool {
	if "" == c.container {
		c.Err = fmt.Errorf("container name must be set")

		return false
	}

	return true
}

func (c *GetContainerShell) Run() error {
	return docker.GetShell(c.container)
}

func (c *GetContainerShell) GetResult() interface{} {
	return nil
}

func CreateGetContainerShellCommand(container string) (*GetContainerShell, error) {
	command := &GetContainerShell{container: container}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
