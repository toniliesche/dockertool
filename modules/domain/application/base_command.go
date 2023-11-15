package application

type BaseCommand struct {
	Err error
}

func (c *BaseCommand) GetError() error {
	return c.Err
}
