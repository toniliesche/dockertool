package images

import (
	"encoding/json"
	"fmt"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"strings"
)

func FetchImage(id string) (*Image, error) {
	images, err := FetchImageList()
	if nil != err {
		return nil, err
	}

	for _, image := range images {
		if id == image.ID {
			return image, nil
		}
	}

	return nil, fmt.Errorf("could not find image %s", id)
}

func FetchImageList() ([]*Image, error) {
	args := []string{
		"image",
		"ls",
		"--format",
		"json",
	}

	output, err := docker.CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	images := make([]*Image, 0, len(output))
	var ls ListResult

	for _, line := range output {
		data := []byte(line)
		if !json.Valid(data) {
			continue
		}

		err = json.Unmarshal(data, &ls)

		images = append(images, &Image{Repository: ls.Repository, Tag: ls.Tag, ID: ls.ID, Size: ls.Size})
	}

	return images, nil
}

func InspectImage(name string) (*InspectResult, error) {
	output, err := docker.CaptureDockerCommand([]string{
		"image",
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
