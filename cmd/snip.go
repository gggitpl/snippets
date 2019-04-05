package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var SnipCommand = &cobra.Command{
	Use:   "snip",
	Short: "List the code snippets saved in snippets, for execution",
	Long:  `List the code snippets saved in snippets and optionally execute`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	SnipCommand.AddCommand(newCommand, gistCommand, editCommand)
	if err := SnipCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
