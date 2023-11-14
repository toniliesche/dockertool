package shared

import (
	"fmt"
	"github.com/toniliesche/dockertool/modules/console"
	"github.com/toniliesche/dockertool/modules/state"
	"runtime"
)

func ShowVersion() {
	fmt.Printf("Author          : %s\n", state.AppState.AuthorName)
	fmt.Printf("E-Mail          : %s\n", state.AppState.AuthorMail)
	fmt.Println()
	fmt.Printf("Build Version   : %s\n", state.AppState.BuildVersion)
	fmt.Printf("Build Commit    : %s\n", state.AppState.BuildCommit)
	fmt.Printf("Build Date      : %s\n", state.AppState.BuildDate)
	fmt.Println()
	fmt.Printf("Build OS        : %s\n", runtime.GOOS)
	fmt.Printf("Build Arch      : %s\n", runtime.GOARCH)
	fmt.Printf("Go Version      : %s\n", runtime.Version())
	fmt.Println()
	fmt.Printf("Release Channel : %s\n", state.AppState.ReleaseChannel)
	fmt.Printf("Use sudo        : %s\n", console.BoolToYesNo(state.AppState.Sudo))
	if state.AppState.Sudo {
		fmt.Printf("User            : %s\n", state.AppState.SudoUser)
	}
	fmt.Println()
}
