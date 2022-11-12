package cmd

import (
	"github.com/spf13/cobra"

	"bee/src/build"
)

var buildCmd = &cobra.Command{
    Use:   "build <path-to-the-function-folder>",
    Short:  "Builds the docker image for the function.",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        build.Build(args[0])
    },
}

func init() {
    rootCmd.AddCommand(buildCmd)
}