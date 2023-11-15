package commands

import "github.com/toniliesche/dockertool/modules/domain/application"

type Runner struct {
}

func (r *Runner) CreateRunCommand(command application.CommandInterface, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	err = command.Run()
	if err != nil {
		return nil, err
	}

	return command.GetResult(), nil
}
