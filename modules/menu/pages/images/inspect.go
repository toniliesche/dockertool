package images

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
	Image string
}

func (p *Inspect) GetHeadline() string {
	return "Show Info"
}

func (p *Inspect) Run() (menu.PageInterface, int, error) {
	data, err := docker.InspectImage(p.Image)
	if err != nil {
		return p.HandleError(err, true)
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

	console.WaitForReturn()

	return nil, 0, nil
}
