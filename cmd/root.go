package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gofig",
	Short: "project sessionizer tool for terminal and tmux environments.",
	Long: `
	gofig is a project sessionizer tool that allows you to specify config files, or a config directory,
	to be used to create tmux sessions, load environments, connect to ssh servers, and more with a single
	command instead of requiring you to manually create the tmux sessions, windows, and panes.

	This also allows you to provide a path to remote or local project templates to quickly spin up
	new go projects basically just wrapping the cool new gonew functionality.
	Read more about that here: https://go.dev/blog/gonew.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
