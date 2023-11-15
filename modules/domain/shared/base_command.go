package shared

type BaseTask struct {
	Err error
}

func (t *BaseTask) GetError() error {
	return t.Err
}
