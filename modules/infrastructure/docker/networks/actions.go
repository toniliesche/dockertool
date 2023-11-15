package networks

import (
	"encoding/json"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"strings"
)

func FetchNetwork(name string) (*Network, error) {
	networks, err := FetchNetworkList()
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

func FetchNetworkList() ([]*Network, error) {
	args := []string{
		"network",
		"ls",
		"--format",
		"json",
	}

	output, err := docker.CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	networks := make([]*Network, 0, len(output))
	var ls ListResult

	for _, line := range output {
		data := []byte(line)
		if !json.Valid(data) {
			continue
		}

		err = json.Unmarshal(data, &ls)

		networks = append(networks, &Network{ID: ls.ID, Name: ls.Name, Driver: ls.Driver, Scope: ls.Scope})
	}

	return networks, nil
}

func InspectNetwork(name string) (*InspectResult, error) {
	output, err := docker.CaptureDockerCommand([]string{
		"network",
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
