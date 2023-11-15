package networks

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/tasks/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/networks"
	"sort"
	"strings"
)

type InspectNetworkTask struct {
	base.Task
	network string
}

func (t *InspectNetworkTask) Validate() bool {
	if "" == t.network {
		t.Err = fmt.Errorf("network name must be set")

		return false
	}

	return true
}

func (t *InspectNetworkTask) Run() error {
	data, err := networks.InspectNetwork(t.network)
	if nil != err {
		return err
	}

	console.PrintHeadline("Network")
	fmt.Printf("Network : %s\n", data.Name)
	fmt.Printf("Driver  : %s\n", data.Driver)
	fmt.Printf("Scope   : %s\n", data.Scope)
	fmt.Println()

	console.PrintHeadline("IP Address Management")
	fmt.Printf("Driver : %s\n", data.IPAM.Driver)
	fmt.Println()
	if len(data.IPAM.Config) > 0 {
		console.PrintHeadline("IP Ranges")
		for _, iprange := range data.IPAM.Config {
			fmt.Printf("Subnet   : %s\n", iprange.Subnet)
			fmt.Printf("Gateway  : %s\n", iprange.Gateway)
			fmt.Printf("IP Range : %s\n", iprange.IPRange)
			fmt.Println()
		}
	}

	console.PrintHeadline("Configuration")
	fmt.Printf("Internal   : %s\n", console.BoolToYesNo(data.Internal))
	fmt.Printf("Attachable : %s\n", console.BoolToYesNo(data.Attachable))
	fmt.Printf("Ingress    : %s\n", console.BoolToYesNo(data.Ingress))
	fmt.Println()

	if len(data.Containers) > 0 {
		console.PrintHeadline("Containers")
		for _, container := range data.Containers {
			fmt.Printf("Name       : %s\n", container.Name)
			fmt.Printf("IP Address : %s\n", container.IPv4Address)
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

func CreateInspectNetworkCommand(network string) (*InspectNetworkTask, error) {
	command := &InspectNetworkTask{network: network}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
