package docker

import (
	"github.com/toniliesche/dockertool/modules/application"
	"github.com/toniliesche/dockertool/modules/infrastructure/shell"
	"os/exec"
)

func RunDockerCommand(args []string, interactive bool, suppressOutput bool) error {
	return shell.RunShell(createCommand(args), interactive, suppressOutput)
}

func CaptureDockerCommand(args []string) ([]string, error) {
	return shell.CaptureShell(createCommand(args))
}

func createCommand(args []string) *exec.Cmd {
	var cmd string
	if application.AppState.Sudo {
		cmd = "sudo"

		if "" != application.AppState.SudoUser {
			args = append([]string{"-u", application.AppState.SudoUser, "docker"}, args...)
		} else {
			args = append([]string{"docker"}, args...)
		}
	} else {
		cmd = "docker"
	}

	return exec.Command(cmd, args...)
}
