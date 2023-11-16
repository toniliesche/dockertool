package compose

import (
	"encoding/json"
	"fmt"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"regexp"
	"strconv"
	"strings"
)

func FetchComposition(name string) (*Composition, error) {
	compositionList, err := FetchCompositionList()
	if err != nil {
		return nil, err
	}

	for _, composition := range compositionList {
		if name == composition.Name {
			return composition, nil
		}
	}

	return nil, fmt.Errorf("could not find composition %s", name)
}

func FetchCompositionList() ([]*Composition, error) {
	args := []string{
		"compose",
		"ls",
		"-a",
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

	runningExp := regexp.MustCompile(`running\((?P<number>\d)\)`)
	var running int
	exitedExp := regexp.MustCompile(`exited\((?P<number>\d)\)`)
	var exited int
	for _, compose := range ls {
		runningMatch := runningExp.FindStringSubmatch(compose.Status)
		if len(runningMatch) > 0 {
			for i, name := range runningExp.SubexpNames() {
				if name == "number" {
					running, _ = strconv.Atoi(runningMatch[i])
				}
			}
		} else {
			running = 0
		}

		exitedMatch := exitedExp.FindStringSubmatch(compose.Status)
		if len(exitedMatch) > 0 {
			for i, name := range exitedExp.SubexpNames() {
				if name == "number" {
					exited, _ = strconv.Atoi(exitedMatch[i])
				}
			}
		} else {
			exited = 0
		}

		compositions = append(compositions, &Composition{Name: compose.Name, ConfigFiles: strings.Split(compose.ConfigFiles, ","), Running: running, Exited: exited})
	}

	return compositions, nil
}
