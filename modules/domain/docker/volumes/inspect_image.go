package volumes

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/application"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"sort"
	"strings"
)

type InspectVolume struct {
	application.BaseCommand
	volume string
}

func (c *InspectVolume) Validate() bool {
	if "" == c.volume {
		c.Err = fmt.Errorf("volume name must be set")

		return false
	}

	return true
}

func (c *InspectVolume) Run() error {
	data, err := docker.InspectVolume(c.volume)
	if err != nil {
		return err
	}

	containers, err := docker.FetchContainersByVolume(c.volume)
	if err != nil {
		return err
	}

	console.PrintHeadline("Volume")
	fmt.Printf("Volume : %s\n", data.Name)
	fmt.Printf("Driver : %s\n", data.Driver)
	fmt.Printf("Scope  : %s\n", data.Scope)
	fmt.Println()

	if len(containers) > 0 {
		console.PrintHeadline("Linked Containers")
		for _, container := range containers {
			fmt.Printf("Name        : %s", strings.TrimSpace(container.Name))

			inspect, err := docker.InspectContainer(container.Name)
			if err != nil {
				return err
			}

			for _, mount := range inspect.Mounts {
				if mount.Name == data.Name {
					fmt.Printf("Mount Point : %s\n", mount.Destination)
					break
				}
			}

			fmt.Println()
		}
	}

	if len(data.Options) > 0 {
		first := true
		keys := make([]string, 0)
		for key := range data.Options {
			if strings.Contains(key, "com.docker.compose") || strings.Contains(key, "org.opencontainers.image") {
				continue
			}

			keys = append(keys, key)
		}

		sort.Strings(keys)
		for _, key := range keys {
			if first {
				console.PrintHeadline("Options")
			}

			fmt.Printf("%s: %s\n", key, data.Options[key])
			first = false
		}

		if !first {
			fmt.Println()
		}
	}

	if len(data.Labels) > 0 {
		first := true
		keys := make([]string, 0)
		for key := range data.Labels {
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

			fmt.Printf("%s: %s\n", key, data.Labels[key])
			first = false
		}

		if !first {
			fmt.Println()
		}
	}

	return nil
}

func (c *InspectVolume) GetResult() interface{} {
	return nil
}

func CreateInspectVolumeCommand(volume string) (*InspectVolume, error) {
	command := &InspectVolume{volume: volume}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
