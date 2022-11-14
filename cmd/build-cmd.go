package cmd

import (
	"github.com/spf13/cobra"

	"bee/src/build"
	"bee/utils"
)

var buildCmd = &cobra.Command{
    Use:   "build <path-to-the-function-folder>",
    Short:  "Builds the docker image for the function.",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        path := utils.Format_path(args[0])
        build.Build(path)
    },
}

func init() {
    rootCmd.AddCommand(buildCmd)
}