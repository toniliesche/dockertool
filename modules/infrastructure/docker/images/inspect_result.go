package images

type InspectResult struct {
	Id              string   `json:"Id"`
	RepoTags        []string `json:"RepoTags"`
	RepoDigests     []string `json:"RepoDigests"`
	Parent          string   `json:"Parent"`
	Comment         string   `json:"Comment"`
	Created         string   `json:"Created"`
	Container       string   `json:"container"`
	ContainerConfig struct {
		Hostname     string            `json:"Hostname"`
		Domainname   string            `json:"Domainname"`
		User         string            `json:"User"`
		AttachStdin  bool              `json:"AttachStdin"`
		AttachStdout bool              `json:"AttachStdout"`
		AttachStderr bool              `json:"AttachStderr"`
		Tty          bool              `json:"Tty"`
		OpenStdin    bool              `json:"OpenStdin"`
		StdinOnce    bool              `json:"StdinOnce"`
		Env          []string          `json:"Env"`
		Cmd          []string          `json:"Cmd"`
		Image        string            `json:"Image"`
		Volumes      interface{}       `json:"Volumes"`
		WorkingDir   string            `json:"WorkingDir"`
		Entrypoint   []string          `json:"Entrypoint"`
		OnBuild      interface{}       `json:"OnBuild"`
		Labels       map[string]string `json:"Labels"`
	} `json:"ContainerConfig"`
	DockerVersion string `json:"DockerVersion"`
	Author        string `json:"Author"`
	Config        struct {
		Hostname     string            `json:"Hostname"`
		Domainname   string            `json:"Domainname"`
		User         string            `json:"User"`
		AttachStdin  bool              `json:"AttachStdin"`
		AttachStdout bool              `json:"AttachStdout"`
		AttachStderr bool              `json:"AttachStderr"`
		Tty          bool              `json:"Tty"`
		OpenStdin    bool              `json:"OpenStdin"`
		StdinOnce    bool              `json:"StdinOnce"`
		Env          []string          `json:"Env"`
		Cmd          []string          `json:"Cmd"`
		Image        string            `json:"Image"`
		Volumes      interface{}       `json:"Volumes"`
		WorkingDir   string            `json:"WorkingDir"`
		Entrypoint   []string          `json:"Entrypoint"`
		OnBuild      interface{}       `json:"OnBuild"`
		Labels       map[string]string `json:"Labels"`
	} `json:"Config"`
	Architecture string `json:"Architecture"`
	Os           string `json:"Os"`
	Size         int    `json:"Size"`
	VirtualSize  int    `json:"VirtualSize"`
	GraphDriver  struct {
		Data struct {
			MergedDir string `json:"MergedDir"`
			UpperDir  string `json:"UpperDir"`
			WorkDir   string `json:"WorkDir"`
		} `json:"Data"`
		Name string
	} `json:"GraphDriver"`
	RootFS struct {
		Type   string   `json:"Type"`
		Layers []string `json:"Layers"`
	} `json:"RootFS"`
	Metadata map[string]string `json:"Metadata"`
}
