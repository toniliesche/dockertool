package base

type Task struct {
	Err    error
	Result interface{}
}

func (t *Task) GetError() error {
	return t.Err
}

func (t *Task) GetResult() interface{} {
	return t.Result
}

func (t *Task) Validate() bool {
	return true
}
