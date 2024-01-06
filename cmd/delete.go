package cmd

import (
	"log"

	"github.com/masonictemple4/gofig/internal/tmux"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [layout.ext]",
	Short: "Delete a resource",
	Long:  `Delete a resource by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No layout specified")
		}

		processDelete(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func processDelete(layout string) {
	if err := tmux.DeleteLayout(layout); err != nil {
		log.Fatal(err)
	}
}
