package networks

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/domain/application"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"sort"
	"strings"
)

type InspectNetwork struct {
	application.BaseCommand
	network string
}

func (c *InspectNetwork) Validate() bool {
	if "" == c.network {
		c.Err = fmt.Errorf("network name must be set")

		return false
	}

	return true
}

func (c *InspectNetwork) Run() error {
	data, err := docker.InspectNetwork(c.network)
	if err != nil {
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

func (c *InspectNetwork) GetResult() interface{} {
	return nil
}

func CreateInspectNetworkCommand(network string) (*InspectNetwork, error) {
	command := &InspectNetwork{network: network}
	if !command.Validate() {
		return nil, command.GetError()
	}

	return command, nil
}
