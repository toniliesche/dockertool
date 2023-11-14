package docker

import (
	"encoding/json"
	"strings"
)

type Network struct {
	Name   string
	ID     string
	Driver string
	Scope  string
}

func GetNetwork(name string) (*Network, error) {
	networks, err := FetchNetworks()
	if err != nil {
		return nil, err
	}

	for _, network := range networks {
		if name == network.Name {
			return network, nil
		}
	}

	return nil, nil
}

func InspectNetwork(name string) (*NetworkInspect, error) {
	output, err := CaptureDockerCommand([]string{"inspect", name})
	if err != nil {
		return nil, err
	}

	jsonString := strings.Join(output, "")

	var inspect = []*NetworkInspect{}
	err = json.Unmarshal([]byte(jsonString), &inspect)
	if err != nil {
		return nil, err
	}

	return inspect[0], nil
}

func FetchNetworks() ([]*Network, error) {
	args := []string{
		"network",
		"ls",
		"--format",
		"{{.ID}} {{.Name}} {{.Driver}} {{.Scope}}",
	}

	output, err := CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	networks := make([]*Network, 0, len(output))

	for _, line := range output {
		parts := strings.Split(line, " ")
		if 4 > len(parts) {
			continue
		}

		networks = append(networks, &Network{ID: parts[0], Name: parts[1], Driver: parts[2], Scope: parts[3]})
	}

	return networks, nil
}
