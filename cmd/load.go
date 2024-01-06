package cmd

import (
	"log"

	"github.com/masonictemple4/gofig/internal/tmux"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load [configFileName.ext]",
	Short: "Load a tmux session layout by name..",
	Long: `Loads a tmux session by the layout name
	you do not need to provide a full file path we will set the installation location
	in the gofig application.
	Loading is very similar to exporting provide the config.yaml or config.json name.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("You must provide a layout name to load")
		}

		processLoad(args[0])

	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}

func processLoad(filename string) {

	tmux.LoadLayout(filename)
}
