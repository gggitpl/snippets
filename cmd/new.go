package cmd

import (
	"github.com/spf13/cobra"
)

var newCommand = &cobra.Command{
	Use:   "new",
	Short: "New command fragment",
	Long:  `Use this command to add a new code version to snippets`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
