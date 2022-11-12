package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:  "Bee",
    Short: "Bee - a simple serverless functions framework, based on kubernetes",
    Long: `Bee is an elite serverless functions framework (not kidding).
One can use bee to create simple functions and deploy them directy in their k8s cluster.`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}