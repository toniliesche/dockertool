package volumes

import (
	"encoding/json"
	"fmt"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"strings"
)

func FetchVolume(name string) (*Volume, error) {
	volumes, err := FetchVolumeList()
	if nil != err {
		return nil, err
	}

	for _, volume := range volumes {
		if name == volume.Name {
			return volume, nil
		}
	}

	return nil, fmt.Errorf("could not find volume %s", name)
}

func FetchVolumeList() ([]*Volume, error) {
	args := []string{
		"volume",
		"ls",
		"--format",
		"json",
	}

	output, err := docker.CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	volumes := make([]*Volume, 0, len(output))
	var ls ListResult

	for _, line := range output {
		data := []byte(line)
		if !json.Valid(data) {
			continue
		}

		err = json.Unmarshal(data, &ls)

		volumes = append(volumes, &Volume{Name: ls.Name, Scope: ls.Scope})
	}

	return volumes, nil
}

func InspectVolume(name string) (*InspectResult, error) {
	output, err := docker.CaptureDockerCommand([]string{
		"volume",
		"inspect",
		name,
	})
	if nil != err {
		return nil, err
	}

	jsonString := strings.Join(output, "")

	var inspect = []*InspectResult{}
	err = json.Unmarshal([]byte(jsonString), &inspect)
	if nil != err {
		return nil, err
	}

	return inspect[0], nil
}
