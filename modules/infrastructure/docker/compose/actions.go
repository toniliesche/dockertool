package compose

import (
	"encoding/json"
	"fmt"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"strings"
)

func FetchCompositionList() ([]*Composition, error) {
	args := []string{
		"compose",
		"ls",
		"--format",
		"json",
	}

	output, err := docker.CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	compositions := make([]*Composition, 0, len(output))
	var ls []*ListResult

	jsonStr := strings.TrimSpace(strings.Join(output, "\n"))

	data := []byte(jsonStr)
	if !json.Valid(data) {
		return nil, fmt.Errorf("docker compose ls did not return valid json")
	}

	err = json.Unmarshal(data, &ls)

	for _, compose := range ls {
		compositions = append(compositions, &Composition{Name: compose.Name})
	}

	return compositions, nil
}
