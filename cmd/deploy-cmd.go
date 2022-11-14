package cmd

import (
	"github.com/spf13/cobra"

	"bee/src/deploy"
	"bee/utils"
)

var deployCmd = &cobra.Command{
    Use:   "deploy <path-to-the-function-folder>",
    Short:  "Deploys the function in the kubernetes cluster.",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        path := utils.Format_path(args[0])
        deploy.Deploy(path)
    },
}

func init() {
    rootCmd.AddCommand(deployCmd)
}