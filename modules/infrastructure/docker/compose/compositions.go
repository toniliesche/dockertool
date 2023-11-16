package compose

type Composition struct {
	Name        string
	ConfigFiles []string
	Running     int
	Exited      int
}

func (c *Composition) IsFullyRunning() bool {
	return c.Exited == 0 && c.Running > 0
}

func (c *Composition) IsFullyStopped() bool {
	return c.Running == 0
}
