package menu

type PageInterface interface {
	RegisterHook()
	HasHook() bool
	GetHeadline() string
	Run() (PageInterface, int, error)
	SetArguments(args []string)
	DontPushToStack() bool
}
