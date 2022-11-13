package cmd

import (
	"github.com/spf13/cobra"

	"bee/src/deploy"
)

var deployCmd = &cobra.Command{
    Use:   "deploy <path-to-the-function-folder>",
    Short:  "Deploys the function in the kubernetes cluster.",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        deploy.Deploy(args[0])
    },
}

func init() {
    rootCmd.AddCommand(deployCmd)
}