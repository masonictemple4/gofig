package tmux

import (
	"encoding/json"
	"testing"
)

func TestWindowsFromString(t *testing.T) {

	sessionName := GetCurrentSessionName()

	args := []string{
		"list-windows",
		"-t",
		sessionName,
		"-F",
		WINDOW_FORMAT,
	}

	output, err := Exec(args)
	if err != nil {
		panic(err)
	}

	windows := WindowsFromString(output)

	data, _ := json.MarshalIndent(windows, "", "  ")
	t.Logf("Windows:\n%s\n", string(data))

}
