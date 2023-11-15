package commands

import "github.com/toniliesche/dockertool/modules/domain/shared"

type TaskRunner struct {
}

func (r *TaskRunner) CreateAndRunTask(command shared.TaskInterface, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	err = command.Run()
	if err != nil {
		return nil, err
	}

	return command.GetResult(), nil
}
