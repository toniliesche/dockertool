package compose

type Composition struct {
	Name        string
	ConfigFiles []string
	Running     int
	Exited      int
}
