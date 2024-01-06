package tmux

import "testing"

func TestExportLayout(t *testing.T) {

	if err := ExportLayout(FormatYAML); err != nil {
		t.Error(err)
	}

	if err := ExportLayout(FormatJSON); err != nil {
		t.Error(err)
	}

}

func TestLoadLayout(t *testing.T) {
}
