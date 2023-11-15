package volumes

type InspectResult struct {
	Name       string            `json:"Name"`
	CreatedAt  string            `json:"CreatedAt"`
	Driver     string            `json:"Driver"`
	Scope      string            `json:"Scope"`
	Mountpoint string            `json:"Mountpoint"`
	Options    map[string]string `json:"Options"`
	Labels     map[string]string `json:"Labels"`
}
