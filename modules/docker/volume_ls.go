package docker

type VolumeLS struct {
	Availability string `json:"Availability"`
	Driver       string `json:"Driver"`
	Group        string `json:"Group"`
	Labels       string `json:"Labels"`
	Links        string `json:"Links"`
	Mountpoint   string `json:"Mountpoint"`
	Name         string `json:"Name"`
	Scope        string `json:"Scope"`
	Size         string `json:"Size"`
	Status       string `json:"Status"`
}
