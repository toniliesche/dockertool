package docker

type NetworkInspect struct {
	Name    string `json:"Name"`
	Created string `json:"Created"`
	Scope   string `json:"Scope"`
	Driver  string `json:"Driver"`
	IPAM    struct {
		Driver  string            `json:"Driver"`
		Options map[string]string `json:"Options"`
		Config  []struct {
			Subnet  string `json:"Subnet"`
			IPRange string `json:"IPRange"`
			Gateway string `json:"Gateway"`
		} `json:"Config"`
	} `json:"IPAM"`
	Internal   bool `json:"Internal"`
	Attachable bool `json:"Attachable"`
	Ingress    bool `json:"Ingress"`
	ConfigOnly bool `json:"ConfigOnly"`
	Containers map[string]struct {
		Name        string `json:"Name"`
		EndpointID  string `json:"EndpointID"`
		MacAddress  string `json:"MacAddress"`
		IPv4Address string `json:"IPv4Address"`
	}
	Options map[string]string `json:"Options"`
	Labels  map[string]string `json:"Labels"`
}
