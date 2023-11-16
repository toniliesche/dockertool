package containers

type InspectResult struct {
	ID      string   `json:"ID"`
	Created string   `json:"Created"`
	Path    string   `json:"Path"`
	Args    []string `json:"Args"`
	State   struct {
		Status     string `json:"Status"`
		Running    bool   `json:"Running"`
		Paused     bool   `json:"Paused"`
		Restarting bool   `json:"Restarting"`
		OOMKilled  bool   `json:"OOMKilled"`
		Dead       bool   `json:"Dead"`
		PID        int    `json:"Pid"`
		ExitCode   int    `json:"ExitCode"`
		StartedAt  string `json:"StartedAt"`
		FinishedAt string `json:"FinishedAt"`
		Health     *struct {
			Status        string `json:"Status"`
			FailingStreak int    `json:"FailingStreak"`
			Log           []struct {
				Start    string `json:"Start"`
				End      string `json:"End"`
				ExitCode int    `json:"ExitCode"`
				Output   string `json:"Output"`
			} `json:"Log"`
		} `json:"Health"`
	} `json:"State"`
	Image        string `json:"Image"`
	Name         string `json:"Name"`
	RestartCount int    `json:"RestartCount"`
	Platform     string `json:"Platform"`
	HostConfig   struct {
		Binds         []string `json:"Binds"`
		NetworkMode   string   `json:"NetworkMode"`
		RestartPolicy struct {
			Name              string `json:"Name"`
			MaximumRetryCount int    `json:"MaximumRetryCount"`
		}
		AutoRemove      bool     `json:"AutoRemove"`
		VolumeDriver    string   `json:"VolumeDriver"`
		VolumesFrom     []string `json:"VolumesFrom"`
		CapAdd          []string `json:"CapAdd"`
		CapDrop         []string `json:"CapDrop"`
		CgroupnsMode    string   `json:"CgroupnsMode"`
		Dns             []string `json:"Dns"`
		DnsOptions      []string `json:"DnsOptions"`
		DnsSearch       []string `json:"DnsSearch"`
		ExtraHosts      []string `json:"ExtraHosts"`
		GroupAdd        []string `json:"GroupAdd"`
		Privileged      bool     `json:"Privileged"`
		PublishAllPorts bool     `json:"PublishAllPorts"`
		ReadonlyRootfs  bool     `json:"ReadonlyRootfs"`
	} `json:"HostConfig"`
	Mounts []struct {
		Name        string `json:"Name"`
		Type        string `json:"Type"`
		Source      string `json:"Source"`
		Destination string `json:"Destination"`
		Mode        string `json:"Mode"`
		RW          bool   `json:"RW"`
		Propagation string `json:"Propagation"`
	} `json:"Mounts"`
	Config struct {
		Hostname     string              `json:"Hostname"`
		Domainname   string              `json:"Domainname"`
		User         string              `json:"User"`
		AttachStdin  bool                `json:"AttachStdin"`
		AttachStdout bool                `json:"AttachStdout"`
		AttachStderr bool                `json:"AttachStderr"`
		ExposedPorts map[string]struct{} `json:"ExposedPorts"`
		Tty          bool                `json:"Tty"`
		OpenStdin    bool                `json:"OpenStdin"`
		StdinOnce    bool                `json:"StdinOnce"`
		Env          []string            `json:"Env"`
		Cmd          []string            `json:"Cmd"`
		Image        string              `json:"ID"`
		Volumes      map[string]struct{} `json:"Volumes"`
		WorkingDir   string              `json:"WorkingDir"`
		Entrypoint   []string            `json:"Entrypoint"`
		Labels       map[string]string   `json:"Labels"`
		StopSignal   string
	} `json:"Config"`
	NetworkSettings struct {
		Bridge string `json:"Bridge"`
		Ports  map[string][]struct {
			HostIp   string `json:"HostIp"`
			HostPort string `json:"HostPort"`
		} `json:"Ports"`
		Networks map[string]struct {
			IPAMConfig  struct{} `json:"IPAMConfig"`
			Links       []string `json:"Links"`
			Aliases     []string `json:"Aliases"`
			Gateway     string   `json:"Gateway"`
			IPAddress   string   `json:"IPAddress"`
			IPPrefixLen int      `json:"IPPrefixLen"`
			MacAddress  string   `json:"MacAddress"`
		} `json:"Networks"`
	} `json:"NetworkSettings"`
}
