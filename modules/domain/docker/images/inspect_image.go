package images

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/application"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"sort"
	"strings"
)

type InspectImage struct {
	application.BaseCommand
	image string
}

func (c *InspectImage) Validate() bool {
	if "" == c.image {
		c.Err = fmt.Errorf("image name must be set")

		return false
	}

	return true
}

func (c *InspectImage) Run() error {
	data, err := docker.InspectImage(c.image)
	if err != nil {
		return err
	}

	console.PrintHeadline("Image")
	fmt.Printf("ID               : %s\n", data.Id)
	fmt.Printf("Architecture     : %s\n", data.Architecture)
	fmt.Printf("Operating System : %s\n", data.Os)
	fmt.Println()

	if len(data.RepoTags) > 0 {
		console.PrintHeadline("Tags")
		for idx, tag := range data.RepoTags {
			fmt.Printf("#%02d : %s\n", idx+1, tag)
		}
		fmt.Println()
	}

	console.PrintHeadline("Execution")
	fmt.Printf("Command    : %s\n", strings.Join(data.Config.Cmd, " "))
	fmt.Printf("Entrypoint : %s\n", strings.Join(data.Config.Entrypoint, " "))
	fmt.Println()

	if len(data.Config.Labels) > 0 {
		first := true
		keys := make([]string, 0)
		for key := range data.Config.Labels {
			if strings.Contains(key, "com.docker.compose") || strings.Contains(key, "org.opencontainers.image") {
				continue
			}

			keys = append(keys, key)
		}

		sort.Strings(keys)
		for _, key := range keys {
			if first {
				console.PrintHeadline("Labels")
			}

			fmt.Printf("%s : %s\n", key, data.Config.Labels[key])
			first = false
		}

		if !first {
			fmt.Println()
		}
	}

	return nil
}

func (c *InspectImage) GetResult() interface{} {
	return nil
}

func CreateInspectImageCommand(image string) (*InspectImage, error) {
	command := &InspectImage{image: image}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
