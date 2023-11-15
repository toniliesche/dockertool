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
	output, err := CaptureDockerCommand([]string{
		"volume",
		"inspect",
		name,
	})
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
		"json",
	}

	output, err := CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	volumes := make([]*Volume, 0, len(output))
	var ls VolumeLS

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
