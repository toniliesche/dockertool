package containers

import (
	"encoding/json"
	"fmt"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"strings"
)

func FetchContainer(name string) (*Container, error) {
	containers, err := FetchContainerList()
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

func FetchContainerList(options ...*FilterOptions) ([]*Container, error) {
	args := []string{
		"container",
		"ls",
		"-a",
		"--format",
		"json",
	}

	return retrieveContainers(args, options...)
}

func FetchContainerListByVolume(volume string) ([]*Container, error) {
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

func GetStdoutFromContainer(name string, follow bool) error {
	container, err := FetchContainer(name)
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

	return docker.RunDockerCommand(args, true, false)
}

func GetShell(name string) error {
	container, err := FetchContainer(name)
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
	err = docker.RunDockerCommand(checkArgs, false, true)
	if err != nil {
		shell = "sh"
	} else {
		shell = "bash"
	}

	return docker.RunDockerCommand([]string{"exec", "-it", container.Name, shell}, true, false)
}

func InspectContainer(name string) (*InspectResult, error) {
	output, err := docker.CaptureDockerCommand([]string{
		"container",
		"inspect",
		name,
	})
	if err != nil {
		return nil, err
	}

	jsonString := strings.Join(output, "")

	var inspect = []*InspectResult{}
	err = json.Unmarshal([]byte(jsonString), &inspect)
	if err != nil {
		return nil, err
	}

	return inspect[0], nil
}

func retrieveContainers(args []string, options ...*FilterOptions) ([]*Container, error) {

	output, err := docker.CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	containers := make([]*Container, 0, len(output))
	var ls ListResult

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