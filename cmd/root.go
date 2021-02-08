package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"k8s.io/component-base/logs"
)

var (
	level          int
	kubeconfigPath string
	namespace      string
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "kuberesource",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logs.GlogSetter(strconv.Itoa(level))
		},
	}

	rootCmd.PersistentFlags().IntVarP(&level, "v", "v", 0, "log level for V logs")
	rootCmd.PersistentFlags().StringVar(&kubeconfigPath, "kubeconfig", "", "The kubeconfig use connect to the Kubernetes cluster.")

	rootCmd.AddCommand(newNodeCommand())

	return rootCmd
}
