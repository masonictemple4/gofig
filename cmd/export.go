package cmd

import (
	"log"
	"path/filepath"

	"github.com/masonictemple4/gofig/internal/tmux"
	"github.com/spf13/cobra"
)

// TODO: Might do config export by filename instead of path
// using the filename as the name of the config etc..
var exportCmd = &cobra.Command{
	Use:   "export [filename]",
	Short: "Export existing session to layout file.",
	Long: `
	Export the existing tmux session to a layout file.
	Note that the filenames should be like "personal.yaml" or "work.json"
	The path is predetermined in the gofig application.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			log.Fatal("No filename specified")
			return
		}

		processExport(args[0])
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

func processExport(filename string) {
	format := filepath.Ext(filename)[1:]

	if !tmux.IsValidOutputFormat(format) {
		log.Fatal("Invalid output format. Please use .json or .yaml")
		return
	}

	err := tmux.ExportLayout(filename)
	if err != nil {
		log.Fatal(err)
	}
}
