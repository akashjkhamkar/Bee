package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"bee/src/create"
)

var createCmd = &cobra.Command{
    Use:   "create <language, eg: python, go> <function-name>",
    Short:  "creates the empty template for the function",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("creating function %s using runtime %s ...\n", args[0], args[1])
        create.Create(args[0], args[1])
    },
}

func init() {
    rootCmd.AddCommand(createCmd)
}