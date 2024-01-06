package cmd

import (
	"github.com/masonictemple4/gofig/internal/tmux"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your layotus",
	Long:  `List all of your layotus`,
	Run: func(cmd *cobra.Command, args []string) {
		tmux.ListLayouts()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
