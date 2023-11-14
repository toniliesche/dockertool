package volumes

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
	"sort"
	"strings"
)

type Inspect struct {
	menu.Base
	Volume string
}

func (p *Inspect) GetHeadline() string {
	return "Show Info"
}

func (p *Inspect) Run() (menu.PageInterface, int, error) {
	data, err := docker.InspectVolume(p.Volume)
	if err != nil {
		return p.HandleError(err, true)
	}

	containers, err := docker.FetchContainersByVolume(p.Volume)
	if err != nil {
		return p.HandleError(err, true)
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
				return p.HandleError(err, true)
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

	console.WaitForReturn()

	return nil, 0, nil
}
