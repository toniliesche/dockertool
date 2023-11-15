package interfaces

type TaskInterface interface {
	Run() error
	Validate() bool
	GetError() error
	GetResult() interface{}
}
