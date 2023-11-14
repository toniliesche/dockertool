package docker

import (
	"github.com/toniliesche/dockertool/modules/cli"
	"github.com/toniliesche/dockertool/modules/state"
	"os/exec"
)

func RunDockerCommand(args []string, interactive bool, suppressOutput bool) error {
	return cli.RunShell(createCommand(args), interactive, suppressOutput)
}

func CaptureDockerCommand(args []string) ([]string, error) {
	return cli.CaptureShell(createCommand(args))
}

func createCommand(args []string) *exec.Cmd {
	var cmd string
	if state.AppState.Sudo {
		cmd = "sudo"

		if "" != state.AppState.SudoUser {
			args = append([]string{"-u", state.AppState.SudoUser, "docker"}, args...)
		} else {
			args = append([]string{"docker"}, args...)
		}
	} else {
		cmd = "docker"
	}

	return exec.Command(cmd, args...)
}
