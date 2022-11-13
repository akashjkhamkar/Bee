package cmd

import (
	"bee/src/push"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
    Use:   "push <function-name> <hub-user>/<repo-name>:<tag>",
    Short:  "Pushes the image to the repository, so that kubernetes can pull it while deploying.",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        push.Push(args[0], args[1])
    },
}

func init() {
    rootCmd.AddCommand(pushCmd)
}