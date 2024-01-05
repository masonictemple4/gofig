package tmux

import "testing"

func TestListPanes(t *testing.T) {
	args := []string{
		"list-panes",
		"-t",
		"1",
		"-F",
		PANE_FORMAT,
	}

	output, err := Exec(args)
	if err != nil {
		t.Fatal(err)
	}

	output = cleanOutput(output)

	t.Logf("Panes:\n%s\n", output)

}
