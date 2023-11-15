package generic

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/application"
	domain_application "github.com/toniliesche/dockertool/modules/domain/shared"
	"github.com/toniliesche/dockertool/modules/infrastructure/console"
	"runtime"
)

type ShowVersionInformationTask struct {
	domain_application.BaseTask
}

func (t *ShowVersionInformationTask) Validate() bool {
	return true
}

func (t *ShowVersionInformationTask) Run() error {
	fmt.Printf("Author          : %s\n", application.AppState.AuthorName)
	fmt.Printf("E-Mail          : %s\n", application.AppState.AuthorMail)
	fmt.Println()
	fmt.Printf("Build Version   : %s\n", application.AppState.BuildVersion)
	fmt.Printf("Build Commit    : %s\n", application.AppState.BuildCommit)
	fmt.Printf("Build Date      : %s\n", application.AppState.BuildDate)
	fmt.Println()
	fmt.Printf("Build OS        : %s\n", runtime.GOOS)
	fmt.Printf("Build Arch      : %s\n", runtime.GOARCH)
	fmt.Printf("Go Version      : %s\n", runtime.Version())
	fmt.Println()
	fmt.Printf("Release Channel : %s\n", application.AppState.ReleaseChannel)
	fmt.Printf("Use sudo        : %s\n", console.BoolToYesNo(application.AppState.Sudo))
	if application.AppState.Sudo {
		fmt.Printf("User            : %s\n", application.AppState.SudoUser)
	}
	fmt.Println()

	return nil
}

func (t *ShowVersionInformationTask) GetResult() interface{} {
	return nil
}

func CreateShowVersionInformationCommand() (*ShowVersionInformationTask, error) {
	return &ShowVersionInformationTask{}, nil
}
