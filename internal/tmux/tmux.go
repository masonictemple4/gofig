package tmux

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	FormatJSON = "json"
	FormatYAML = "yaml"
)

func IsValidOutputFormat(format string) bool {
	return format == FormatJSON || format == FormatYAML
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
	case FormatJSON:
		if err := json.NewEncoder(outFile).Encode(sessions); err != nil {
			panic(err)
		}
	case FormatYAML:
		if err := yaml.NewEncoder(outFile).Encode(sessions); err != nil {
			panic(err)
		}
	}

	println("Exported layout to layout." + format)

}

// Loads a new tmux server with the layout.
func LoadLayout(path string) {
	// Use the execandreplace command here

	inFmt := filepath.Ext(path)

	inFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var sessions []Session

	switch inFmt {
	case FormatJSON:
		if err := json.Unmarshal(inFile, &sessions); err != nil {
			panic(err)
		}
	case FormatYAML:
		if err := yaml.Unmarshal(inFile, &sessions); err != nil {
			panic(err)
		}
	}

}
