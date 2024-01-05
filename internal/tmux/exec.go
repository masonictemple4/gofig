package tmux

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const LINE_DELIMITER = "\n"
const FIELD_DELIMITER = "|"

// Run tmux commands.
// Note: Calls cleanoutput to remove the trailing newline by
// default.
func Exec(args []string) (string, error) {
	tmux, err := exec.LookPath("tmux")
	if err != nil {
		return "", err
	}

	result, err := exec.Command(tmux, args...).Output()
	if err != nil {
		return "", err
	}

	return cleanOutput(string(result)), nil
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

/*
	convenience methods for standard handling of
	tmux command outputs.
*/

// Removes empty line at the end of the
// output string that results in an extra item
// when splitting lines.
func cleanOutput(output string) string {
	output = strings.TrimSuffix(output, LINE_DELIMITER)
	return output
}

func splitLines(output string) []string {
	lines := strings.Split(output, LINE_DELIMITER)
	return lines
}

func splitFields(output string) []string {
	fields := strings.Split(output, FIELD_DELIMITER)
	return fields
}
