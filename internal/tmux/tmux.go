package tmux

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	OutputFormatJSON = "json"
	OutputFormatYAML = "yaml"
)

func IsValidOutputFormat(format string) bool {
	return format == OutputFormatJSON || format == OutputFormatYAML
}

// Exports the current layout of the current tmux server.
func ExportLayout(format string) {
	if !IsValidOutputFormat(format) {
		panic("Invalid output format")
	}

	sessions := GetSessionsList()

	outFileName := "layout." + format
	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	switch format {
	case OutputFormatJSON:
		if err := json.NewEncoder(outFile).Encode(sessions); err != nil {
			panic(err)
		}
	case OutputFormatYAML:
		if err := yaml.NewEncoder(outFile).Encode(sessions); err != nil {
			panic(err)
		}
	}

	println("Exported layout to layout." + format)

}

// Loads a new tmux server with the layout.
func LoadLayout(path string) {
}
