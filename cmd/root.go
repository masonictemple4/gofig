package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofig",
	Short: "project sessionizer tool for terminal and tmux environments.",
	Long: `
	gofig is a tmux tool to export and save configurations for tmux sessions locally,
	and loading with a simple command.


	Layouts: Layouts are saved to your ~/.config/tmux/layouts directory.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
