package tmux

import "testing"

func TestExportLayout(t *testing.T) {

	ExportLayout(FormatYAML)

	ExportLayout(FormatJSON)

}

func TestLoadLayout(t *testing.T) {
}
