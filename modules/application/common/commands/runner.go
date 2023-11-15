package commands

import (
	"github.com/toniliesche/dockertool/modules/domain/tasks/interfaces"
)

type TaskRunner struct {
}

func (r *TaskRunner) CreateAndRunTask(command interfaces.TaskInterface, err error) (interface{}, error) {
	if nil != err {
		return nil, err
	}

	err = command.Run()
	if nil != err {
		return nil, err
	}

	return command.GetResult(), nil
}
