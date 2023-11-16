package system

type Info struct {
	ID                string     `json:"ID"`
	Containers        int        `json:"Containers"`
	ContainersRunning int        `json:"ContainersRunning"`
	ContainersPaused  int        `json:"ContainersPaused"`
	ContainersStopped int        `json:"ContainersStopped"`
	Images            int        `json:"Images"`
	Driver            string     `json:"Driver"`
	DriverStatus      [][]string `json:"DriverStatus"`
	Plugins           struct {
		Volume        []string    `json:"Volume"`
		Network       []string    `json:"Network"`
		Authorization interface{} `json:"Authorization"`
		Log           []string    `json:"Log"`
	}
	MemoryLimit        bool   `json:"MemoryLimit"`
	SwapLimit          bool   `json:"SwapLimit"`
	CpuCfsPeriod       bool   `json:"CpuCfsPeriod"`
	CpuCfsQuota        bool   `json:"CpuCfsQuota"`
	CPUShares          bool   `json:"CPUShares"`
	CPUSet             bool   `json:"CPUSet"`
	PidsLimit          bool   `json:"PidsLimit"`
	IPv4Forwarding     bool   `json:"IPv4Forwarding"`
	BridgeNfIptables   bool   `json:"BridgeNfIptables"`
	BridgeNfIP6tables  bool   `json:"BridgeNfIP6Tables"`
	Debug              bool   `json:"Debug"`
	NFd                int    `json:"NFd"`
	OomKillDisable     bool   `json:"OomKillDisable"`
	NGoroutines        int    `json:"NGoroutines"`
	SystemTime         string `json:"SystemTime"`
	LoggingDriver      string `json:"LoggingDriver"`
	CgroupDriver       string `json:"CgroupDriver"`
	CgroupVersion      string `json:"CgroupVersion"`
	NEventsListener    int    `json:"NEventsListener"`
	KernelVersion      string `json:"KernelVersion"`
	OperatingSystem    string `json:"OperatingSystem"`
	OSVersion          string `json:"OSVersion"`
	OSType             string `json:"OSType"`
	Architecture       string `json:"Architecture"`
	IndexServerAddress string `json:"IndexServerAddress"`
	RegistryConfig     struct {
		AllowNondistributableArtifactsCIDRs     interface{} `json:"AllowNondistributableArtifactsCIDRs"`
		AllowNondistributableArtifactsHostnames interface{} `json:"AllowNondistributableArtifactsHostnames"`
		InsecureRegistryCIDRs                   []string    `json:"InsecureRegistryCIDRs"`
		IndexConfigs                            struct {
			DockerIo struct {
				Name     string        `json:"Name"`
				Mirrors  []interface{} `json:"Mirrors"`
				Secure   bool          `json:"Secure"`
				Official bool          `json:"Official"`
			} `json:"docker.io"`
		} `json:"IndexConfigs"`
		Mirrors interface{} `json:"Mirrors"`
	} `json:"RegistryConfig"`
	NCPU              int           `json:"NCPU"`
	MemTotal          int64         `json:"MemTotal"`
	GenericResources  interface{}   `json:"GenericResources"`
	DockerRootDir     string        `json:"DockerRootDir"`
	HttpProxy         string        `json:"HttpProxy"`
	HttpsProxy        string        `json:"HttpsProxy"`
	NoProxy           string        `json:"NoProxy"`
	Name              string        `json:"Name"`
	Labels            []interface{} `json:"Labels"`
	ExperimentalBuild bool          `json:"ExperimentalBuild"`
	ServerVersion     string        `json:"ServerVersion"`
	Runtimes          struct {
		IoContainerdRuncV2 struct {
			Path string `json:"path"`
		} `json:"io.containerd.runc.v2"`
		Runc struct {
			Path string `json:"path"`
		} `json:"runc"`
	} `json:"Runtimes"`
	DefaultRuntime string `json:"DefaultRuntime"`
	Swarm          struct {
		NodeID           string      `json:"NodeID"`
		NodeAddr         string      `json:"NodeAddr"`
		LocalNodeState   string      `json:"LocalNodeState"`
		ControlAvailable bool        `json:"ControlAvailable"`
		Error            string      `json:"Error"`
		RemoteManagers   interface{} `json:"RemoteManagers"`
	} `json:"Swarm"`
	LiveRestoreEnabled bool   `json:"LiveRestoreEnabled"`
	Isolation          string `json:"Isolation"`
	InitBinary         string `json:"InitBinary"`
	ContainerdCommit   struct {
		ID       string `json:"ID"`
		Expected string `json:"Expected"`
	} `json:"ContainerdCommit"`
	RuncCommit struct {
		ID       string `json:"ID"`
		Expected string `json:"Expected"`
	} `json:"RuncCommit"`
	InitCommit struct {
		ID       string `json:"ID"`
		Expected string `json:"Expected"`
	} `json:"InitCommit"`
	SecurityOptions     []string `json:"SecurityOptions"`
	DefaultAddressPools []struct {
		Base string `json:"Base"`
		Size int    `json:"Size"`
	} `json:"DefaultAddressPools"`
	Warnings   interface{} `json:"Warnings"`
	ClientInfo struct {
		Debug    bool `json:"Debug"`
		Platform struct {
			Name string `json:"Name"`
		} `json:"Platform"`
		Version   string `json:"Version"`
		GitCommit string `json:"GitCommit"`
		GoVersion string `json:"GoVersion"`
		Os        string `json:"Os"`
		Arch      string `json:"Arch"`
		BuildTime string `json:"BuildTime"`
		Context   string `json:"Context"`
		Plugins   []struct {
			SchemaVersion    string `json:"SchemaVersion"`
			Vendor           string `json:"Vendor"`
			Version          string `json:"Version"`
			ShortDescription string `json:"ShortDescription"`
			Name             string `json:"Name"`
			Path             string `json:"Path"`
		} `json:"Plugins"`
		Warnings interface{} `json:"Warnings"`
	} `json:"ClientInfo"`
}
