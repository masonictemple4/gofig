package tmux

import (
	"os"
	"os/exec"
	"syscall"
)

// Use the tmux cli to setup environment.
func Exec(args []string) (string, error) {
	tmux, err := exec.LookPath("tmux")
	if err != nil {
		return "", err
	}

	result, err := exec.Command(tmux, args...).Output()
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// Use syscall to replace the go process with the new executed one.
// use this when starting a new environment.
func ExecAndReplace(args []string) error {
	tmux, err := exec.LookPath("tmux")
	if err != nil {
		return err
	}

	args = append([]string{tmux}, args...)

	syscall.Exec(tmux, args, os.Environ())
	if err != nil {
		return err
	}

	return nil
}
