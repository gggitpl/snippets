package cmd

import "github.com/spf13/cobra"

var gistCommand = &cobra.Command{
	Use:   "gist",
	Short: "Sync code snippet",
	Long:  `Configure gist to synchronize code snippets on different machines`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
