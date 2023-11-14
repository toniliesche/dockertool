package docker

import (
	"encoding/json"
	"strings"
)

type Volume struct {
	Name  string
	Scope string
}

func GetVolume(name string) (*Volume, error) {
	volumes, err := FetchVolumes()
	if err != nil {
		return nil, err
	}

	for _, volume := range volumes {
		if name == volume.Name {
			return volume, nil
		}
	}

	return nil, nil
}

func InspectVolume(name string) (*VolumeInspect, error) {
	output, err := CaptureDockerCommand([]string{"inspect", name})
	if err != nil {
		return nil, err
	}

	jsonString := strings.Join(output, "")

	var inspect = []*VolumeInspect{}
	err = json.Unmarshal([]byte(jsonString), &inspect)
	if err != nil {
		return nil, err
	}

	return inspect[0], nil
}

func FetchVolumes() ([]*Volume, error) {
	args := []string{
		"volume",
		"ls",
		"--format",
		"{{.Name}} {{.Scope}}",
	}

	output, err := CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	volumes := make([]*Volume, 0, len(output))

	for _, line := range output {
		parts := strings.Split(line, " ")
		if 2 > len(parts) {
			continue
		}

		volumes = append(volumes, &Volume{Name: parts[0], Scope: parts[1]})
	}

	return volumes, nil
}
