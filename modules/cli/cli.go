package cli

import (
	"bytes"
	"fmt"
	"github.com/toniliesche/dockertool/modules/state"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func RunShell(cmd *exec.Cmd, interactive bool, suppressOutput bool) error {
	return run(cmd, os.Stdout, os.Stdin, interactive, suppressOutput)
}

func CaptureShell(cmd *exec.Cmd) ([]string, error) {
	output := make([]string, 0)

	writer := bytes.NewBufferString("")
	err := run(cmd, writer, writer, false, false)
	if err != nil {
		return output, err
	}

	output = strings.Split(writer.String(), "\n")

	return output, nil
}

func run(cmd *exec.Cmd, outputWriter io.Writer, errorWriter io.Writer, interactive bool, suppressOutput bool) error {
	if !suppressOutput {
		cmd.Stdout = outputWriter
		cmd.Stderr = os.Stderr
	}

	if interactive {
		cmd.Stdin = os.Stdin
	}

	c := make(chan os.Signal)
	defer close(c)
	signal.Reset(os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer state.AppState.ResetHandler()

	go func() {
		<-c
		cmd.Process.Signal(os.Interrupt)
	}()

	err := cmd.Run()
	if !suppressOutput {
		fmt.Println()
		if err != nil {
			fmt.Printf("An error occured: %s\n", err.Error())
		}
	}

	return err
}
