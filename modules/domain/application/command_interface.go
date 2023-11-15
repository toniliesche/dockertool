package application

type CommandInterface interface {
	Run() error
	Validate() bool
	GetError() error
	GetResult() interface{}
}
