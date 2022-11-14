package cmd

import (
	"bee/src/push"
	"bee/utils"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
    Use:   "push <path-to-function-folder> <repo-user>/<repo-name>:<tag>",
    Short:  "Pushes the image to the repository, so that kubernetes can pull it while deploying.",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        path := utils.Format_path(args[0])
        push.Push(path, args[1])
    },
}

func init() {
    rootCmd.AddCommand(pushCmd)
}