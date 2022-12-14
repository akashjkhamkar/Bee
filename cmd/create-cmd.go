package cmd

import (
	"github.com/spf13/cobra"

	"bee/src/create"
)

var createCmd = &cobra.Command{
    Use:   "create <language, eg: python, go> <function-name>",
    Short:  "Creates the empty template for the function",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        create.Create(args[0], args[1])
    },
}

func init() {
    rootCmd.AddCommand(createCmd)
}