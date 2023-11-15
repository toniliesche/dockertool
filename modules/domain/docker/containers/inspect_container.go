package containers

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/application"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"sort"
	"strings"
)

type InspectContainer struct {
	application.BaseCommand
	container string
}

func (c *InspectContainer) Validate() bool {
	if "" == c.container {
		c.Err = fmt.Errorf("container name must be set")

		return false
	}

	return true
}

func (c *InspectContainer) Run() error {
	data, err := docker.InspectContainer(c.container)
	if err != nil {
		return err
	}

	console.PrintHeadline("container")
	fmt.Printf("container  : %s\n", data.Name[1:])
	fmt.Printf("ID         : %s\n", data.Config.Image)
	fmt.Printf("Hostname   : %s\n", data.Config.Hostname)
	fmt.Printf("Domainname : %s\n", data.Config.Domainname)
	fmt.Println()

	console.PrintHeadline("State")
	fmt.Printf("Status : %s\n", data.State.Status)
	fmt.Println()

	if data.State.Running && nil != data.State.Health {
		console.PrintHeadline("Health")
		fmt.Printf("Status         : %s\n", data.State.Health.Status)
		if "healthy" != data.State.Health.Status {
			fmt.Printf("Failing Streak : %d\n", data.State.Health.FailingStreak)

			logLength := len(data.State.Health.Log)
			fmt.Printf("Last Entry     : Code %d, Message %s\n", data.State.Health.Log[logLength-1].ExitCode, strings.Join(strings.Split(data.State.Health.Log[logLength-1].Output, "\n"), ", "))
		}

		fmt.Println()
	}

	console.PrintHeadline("Execution")
	fmt.Printf("Command    : %s\n", strings.Join(data.Config.Cmd, " "))
	fmt.Printf("Entrypoint : %s\n", strings.Join(data.Config.Entrypoint, " "))
	fmt.Printf("Privileged : %s\n", console.BoolToYesNo(data.HostConfig.Privileged))
	fmt.Printf("Open stdin : %s\n", console.BoolToYesNo(data.Config.OpenStdin))
	fmt.Printf("Tty open   : %s\n", console.BoolToYesNo(data.Config.Tty))
	fmt.Println()

	if len(data.Mounts) > 0 {
		console.PrintHeadline("Volumes")
		for _, mount := range data.Mounts {
			if "volume" == mount.Type {
				fmt.Printf("Volume      : %s\n", mount.Name)
			} else {
				fmt.Printf("Mount       : %s\n", mount.Source)
			}
			fmt.Printf("Destination : %s\n", mount.Destination)
			if mount.RW {
				fmt.Println("Mode        : writable")
			} else {
				fmt.Println("Mode        : read-only")
			}
			fmt.Println()
		}
	}

	if len(data.NetworkSettings.Networks) > 0 {
		console.PrintHeadline("Networks")
		for name, network := range data.NetworkSettings.Networks {
			fmt.Printf("Network : %s\n", name)
			if data.State.Running {
				fmt.Printf("IP      : %s/%d\n", network.IPAddress, network.IPPrefixLen)
			}
			fmt.Printf("Aliases : %s\n", strings.Join(network.Aliases, ", "))
			fmt.Println()
		}
	}

	if len(data.NetworkSettings.Ports) > 0 {
		first := true
		for dstPort, fwdPorts := range data.NetworkSettings.Ports {
			for idx, fwdPort := range fwdPorts {
				if first {
					console.PrintHeadline("Ports")
				}

				fmt.Printf("#%02d : %s:%s->%s\n", idx+1, fwdPort.HostIp, fwdPort.HostPort, dstPort)
				first = false
			}
		}
		if !first {
			fmt.Println()
		}
	}

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

func (c *InspectContainer) GetResult() interface{} {
	return nil
}

func CreateInspectContainerCommand(container string) (*InspectContainer, error) {
	command := &InspectContainer{container: container}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
