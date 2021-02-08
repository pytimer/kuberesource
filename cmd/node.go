package cmd

import (
	"context"
	"os"

	"github.com/pytimer/kuberesource/pkg/kube"
	"github.com/pytimer/kuberesource/pkg/scraper"

	"github.com/spf13/cobra"
	"k8s.io/klog"
)

func newNodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "node",
		Run: func(cmd *cobra.Command, args []string) {
			names := args
			client, err := kube.NewClient(kubeconfigPath)
			if err != nil {
				klog.Error(err)
				os.Exit(1)
			}
			s := scraper.NewScraper(client)
			s.FetchResourceQuota(context.Background(), names)
		},
	}
	return cmd
}
