package main

import (
	"math/rand"
	"time"

	"github.com/pytimer/kuberesource/cmd"

	"k8s.io/component-base/logs"
	"k8s.io/klog"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	logs.InitLogs()
	defer logs.FlushLogs()

	c := cmd.NewCommand()
	if err := c.Execute(); err != nil {
		klog.Fatal(err)
	}
}
