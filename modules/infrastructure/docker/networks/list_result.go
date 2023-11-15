package networks

type ListResult struct {
	CreatedAt string `json:"CreatedAt"`
	Driver    string `json:"Driver"`
	ID        string `json:"ID"`
	IPv6      string `json:"IPv6"`
	Internal  string `json:"Internal"`
	Labels    string `json:"Labels"`
	Name      string `json:"Name"`
	Scope     string `json:"Scope"`
}
