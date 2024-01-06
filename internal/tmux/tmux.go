package tmux

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
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

func generateLayoutDir() {
	home, err := os.UserHomeDir()

	layoutDir := filepath.Join(home, ".config/tmux/layouts")

	if err != nil {
		panic(err)
	}

	if err := os.MkdirAll(layoutDir, 0755); err != nil {
		panic(err)
	}
}

func getOutputPath(filename string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	layoutDir := filepath.Join(home, ".config/tmux/layouts")

	return filepath.Join(layoutDir, filename)
}

// Exports the current layout of the current tmux server.
func ExportLayout(filename string) error {
	format := filepath.Ext(filename)[1:]

	if !IsValidOutputFormat(format) {
		return errors.New("Invalid output format")
	}

	generateLayoutDir()

	outFileName := getOutputPath(filename)

	sessions := GetSessionsList()

	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer outFile.Close()

	switch format {
	case FormatJSON:
		if err := json.NewEncoder(outFile).Encode(sessions); err != nil {
			return err
		}
	case FormatYAML:
		if err := yaml.NewEncoder(outFile).Encode(sessions); err != nil {
			return err
		}
	}

	println("Exported layout to layout." + outFileName)

	return nil
}

// Loads a new tmux server with the layout.
func LoadLayout(filename string) {
	// Use the execandreplace command here
	inFmt := filepath.Ext(filename)[1:]

	path := getOutputPath(filename)

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
	default:
		panic("Invalid format use json or yaml")
	}

	for _, s := range sessions {
		mainWindow := s.Windows[0]

		args := []string{
			"new-session",
			"-d",
			"-n",
			mainWindow.Name,
			"-D",
			"-s",
			s.Name,
			"-c",
			mainWindow.Panes[0].WorkdDir,
		}

		// TODO: Figure out what to do with out here.
		_, err := Exec(args)
		if err != nil {
			panic(err)
		}

		for wid, w := range s.Windows {
			startDir := w.Panes[0].WorkdDir

			if wid > 0 {
				args := []string{
					"new-window",
					"-k",
					"-n",
					w.Name,
					"-t",
					fmt.Sprintf("%d", w.Id),
					"-c",
					startDir,
				}

				_, err := Exec(args)
				if err != nil {
					panic(err)
				}
			}

			for pid := range w.Panes {
				if pid > 0 {

					args := []string{
						"split-window",
						"-t",
						fmt.Sprintf("%d", w.Id),
						"-c",
						w.Panes[pid].WorkdDir,
					}

					_, err := Exec(args)
					if err != nil {
						panic(err)
					}
				}
			}

		}
	}

	args := []string{
		"attach-session",
		"-t",
		fmt.Sprintf("%s:0", sessions[0].Name),
	}

	if err := ExecAndReplace(args); err != nil {
		panic(err)
	}
}

// Will remove the local layout file
func DeleteLayout(filename string) error {
	outPath := getOutputPath(filename)

	if err := os.Remove(outPath); err != nil {
		return err
	}

	println("Successfully removed " + outPath)

	return nil

}

func ListLayouts() error {
	layoutDir := getOutputPath("")

	ls, err := exec.LookPath("ls")
	if err != nil {
		return err
	}

	args := []string{
		"-al",
		layoutDir,
	}

	output, err := exec.Command(ls, args...).Output()
	if err != nil {
		return err
	}

	println(string(output))

	return nil
}
