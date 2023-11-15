package docker

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Container struct {
	Name  string
	ID    string
	State string
	Image string
}

func (c *Container) IsRunning() bool {
	return "running" == c.State
}

func (c *Container) IsRunningString() string {
	if c.IsRunning() {
		return "yes"
	}

	return "no"
}

func GetStdoutFromContainer(name string, follow bool) error {
	container, err := GetContainer(name)
	if err != nil {
		return err
	}

	if container == nil {
		return fmt.Errorf("container not found")
	}

	if !container.IsRunning() {
		return fmt.Errorf("container is not running")
	}

	args := []string{
		"logs",
		"--tail",
		"20",
	}

	if follow {
		args = append(args, "-f")
	}

	args = append(args, container.Name)

	return RunDockerCommand(args, true, false)
}

func GetContainer(name string) (*Container, error) {
	containers, err := FetchContainers()
	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		if name == container.Name {
			return container, nil
		}
	}

	return nil, nil
}

func GetShell(name string) error {
	container, err := GetContainer(name)
	if err != nil {
		return err
	}

	if container == nil {
		return fmt.Errorf("container not found")
	}

	if !container.IsRunning() {
		return fmt.Errorf("container is not running")
	}

	checkArgs := []string{"exec", "-it", container.Name, "which", "bash"}
	var shell string
	err = RunDockerCommand(checkArgs, false, true)
	if err != nil {
		shell = "sh"
	} else {
		shell = "bash"
	}

	return RunDockerCommand([]string{"exec", "-it", container.Name, shell}, true, false)
}

func InspectContainer(name string) (*ContainerInspect, error) {
	output, err := CaptureDockerCommand([]string{
		"container",
		"inspect",
		name,
	})
	if err != nil {
		return nil, err
	}

	jsonString := strings.Join(output, "")

	var inspect = []*ContainerInspect{}
	err = json.Unmarshal([]byte(jsonString), &inspect)
	if err != nil {
		return nil, err
	}

	return inspect[0], nil
}

func FetchContainersByVolume(volume string) ([]*Container, error) {
	args := []string{
		"container",
		"ls",
		"--filter",
		fmt.Sprintf("volume=%s", volume),
		"--format",
		"json",
	}

	return retrieveContainers(args)
}

func FetchContainers(options ...*FilterOptions) ([]*Container, error) {
	args := []string{
		"container",
		"ls",
		"--format",
		"json",
	}

	return retrieveContainers(args, options...)
}

func retrieveContainers(args []string, options ...*FilterOptions) ([]*Container, error) {

	output, err := CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	containers := make([]*Container, 0, len(output))
	var ls ContainerLS

	for _, line := range output {
		data := []byte(line)
		if !json.Valid(data) {
			continue
		}

		err = json.Unmarshal(data, &ls)
		if err != nil {
			return nil, err
		}

		containers = append(containers, &Container{ID: ls.ID, Name: ls.Names, Image: ls.Image, State: ls.State})
	}

	if len(options) > 0 {
		containers, err = runFilters(options[0], containers)
		if err != nil {
			return nil, err
		}
	}

	return containers, nil
}
