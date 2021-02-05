package main

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/pytimer/kuberesource/pkg/kube"
	"github.com/pytimer/kuberesource/pkg/scraper"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/component-base/logs"
	"k8s.io/klog"
)

var (
	level          int
	kubeconfigPath string
)

func main() {
	rand.Seed(time.Now().UnixNano())

	logs.InitLogs()
	defer logs.FlushLogs()

	pflag.IntVarP(&level, "v", "v", 0, "log level for V logs")
	pflag.StringVar(&kubeconfigPath, "kubeconfig", "", "The kubeconfig use connect to the Kubernetes cluster.")

	logs.GlogSetter(strconv.Itoa(level))

	root := cobra.Command{
		Use: "kuberesource",
		Run: func(cmd *cobra.Command, args []string) {
			names := args
			client, err := kube.NewClient(kubeconfigPath)
			if err != nil {
				return
			}
			s := scraper.NewScraper(client)
			s.FetchResourceQuota(context.Background(), names)
		},
	}

	if err := root.Execute(); err != nil {
		klog.Fatal(err)
	}
}
