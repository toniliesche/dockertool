package docker

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Composition struct {
	Name        string
	ConfigFiles []string
	Running     int
	Exited      int
}

func FetchCompositions() ([]*Composition, error) {
	args := []string{
		"compose",
		"ls",
		"--format",
		"json",
	}

	output, err := CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	compositions := make([]*Composition, 0, len(output))
	var ls []*ComposeLS

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
