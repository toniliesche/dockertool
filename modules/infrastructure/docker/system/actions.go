package system

import (
	"encoding/json"
	"fmt"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
)

func GetSystemInfo() (*Info, error) {
	args := []string{
		"system",
		"info",
		"--format",
		"json",
	}

	output, err := docker.CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	var info Info

	for _, line := range output {
		data := []byte(line)

		if !json.Valid(data) {
			continue
		}
		err = json.Unmarshal(data, &info)
		if nil != err {
			return nil, err
		}

		return &info, nil
	}

	return nil, fmt.Errorf("something went wrong while parsing system info")
}

func CheckComposePlugin() error {
	args := []string{
		"compose",
		"version",
	}

	_, err := docker.CaptureDockerCommand(args)
	return err
}
