package cmd

import (
	"github.com/spf13/cobra"
)

var editCommand = &cobra.Command{
	Use:   "edit",
	Short: "Edit the code snippet you have added",
	Long:  `Modify or delete code snippets that have been added to snippets`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
