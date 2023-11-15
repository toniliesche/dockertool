package containers

type Container struct {
	Name  string
	ID    string
	State string
	Image string
}

func (c *Container) IsRunning() bool {
	return "running" == c.State
}

func (c *Container) IsRunningString() string {
	if c.IsRunning() {
		return "yes"
	}

	return "no"
}
