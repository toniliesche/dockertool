package volumes

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/tasks/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/containers"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/volumes"
	"sort"
	"strings"
)

type InspectVolumeTask struct {
	base.Task
	volume string
}

func (t *InspectVolumeTask) Validate() bool {
	if "" == t.volume {
		t.Err = fmt.Errorf("volume name must be set")

		return false
	}

	return true
}

func (t *InspectVolumeTask) Run() error {
	data, err := volumes.InspectVolume(t.volume)
	if nil != err {
		return err
	}

	containerList, err := containers.FetchContainerListByVolume(t.volume)
	if nil != err {
		return err
	}

	console.PrintHeadline("Volume")
	fmt.Printf("Volume : %s\n", data.Name)
	fmt.Printf("Driver : %s\n", data.Driver)
	fmt.Printf("Scope  : %s\n", data.Scope)
	fmt.Println()

	if len(containerList) > 0 {
		console.PrintHeadline("Linked Containers")
		for _, container := range containerList {
			fmt.Printf("Name        : %s", strings.TrimSpace(container.Name))

			inspect, err := containers.InspectContainer(container.Name)
			if nil != err {
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

func CreateInspectVolumeCommand(volume string) (*InspectVolumeTask, error) {
	command := &InspectVolumeTask{volume: volume}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
