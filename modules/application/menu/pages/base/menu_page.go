package base

import (
	"github.com/toniliesche/dockertool/modules/application/common/commands"
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/domain/tasks/generic"
)

type MenuPage struct {
	Page
	commands.TaskRunner
}

func (m *MenuPage) CreateAndRunMenuTask(menuEntries models.EntryList, specialEntries models.EntryList, err error) (interfaces.PageInterface, int, error) {
	if err != nil {
		return m.HandleError(err, true)
	}

	result, err := m.CreateAndRunTask(generic.CreateShowMenuTask(menuEntries, specialEntries))
	if nil != err {
		return m.HandleError(err, true)
	}

	if nil != result {
		return m.EvaluateResult(result.(*models.Entry))
	}

	return m.EvaluateResult(nil)
}
