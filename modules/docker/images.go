package docker

import (
	"encoding/json"
	"strings"
)

type Image struct {
	ID         string
	Repository string
	Tag        string
	Size       string
}

func GetImage(id string) (*Image, error) {
	images, err := FetchImages()
	if err != nil {
		return nil, err
	}

	for _, image := range images {
		if id == image.ID {
			return image, nil
		}
	}

	return nil, nil
}

func InspectImage(name string) (*ImageInspect, error) {
	output, err := CaptureDockerCommand([]string{
		"image",
		"inspect",
		name,
	})
	if err != nil {
		return nil, err
	}

	jsonString := strings.Join(output, "")

	var inspect = []*ImageInspect{}
	err = json.Unmarshal([]byte(jsonString), &inspect)
	if err != nil {
		return nil, err
	}

	return inspect[0], nil
}

func FetchImages() ([]*Image, error) {
	args := []string{
		"image",
		"ls",
		"--format",
		"json",
	}

	output, err := CaptureDockerCommand(args)
	if nil != err {
		return nil, err
	}

	images := make([]*Image, 0, len(output))
	var ls ImageLS

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
