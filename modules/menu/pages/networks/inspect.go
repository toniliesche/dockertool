package networks

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/docker"
	"github.com/toniliesche/dockertool/modules/menu"
)

type Inspect struct {
	menu.Base
	Network string
}

func (p *Inspect) GetHeadline() string {
	return "Show Info"
}

func (p *Inspect) Run() (menu.PageInterface, int, error) {
	data, err := docker.InspectNetwork(p.Network)
	if err != nil {
		return p.HandleError(err, true)
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

	console.WaitForReturn()

	return nil, 0, nil
}
