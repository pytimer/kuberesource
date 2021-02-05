package scraper

import (
	"context"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pytimer/kuberesource/pkg/node"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type scraper struct {
	kubeClient kubernetes.Interface
}

func NewScraper(client kubernetes.Interface) *scraper {
	return &scraper{
		kubeClient: client,
	}
}

/*
	            nodename
+--------+----------+--------+-------------+
| Name   | Requests | Limits | Allocatable |
+--------+----------+--------+-------------+
| cpu    | 832m     | 4300m  | 3600m       |
| memory | 520Mi    | 3694Mi | 6333800Ki   |
+--------+----------+--------+-------------+
*/
func (s *scraper) FetchResourceQuota(ctx context.Context) error {
	nodes, err := s.kubeClient.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, n := range nodes.Items {
		podList, err := node.GetNonTerminatedPodsOfNode(s.kubeClient, n.Name)
		if err != nil {
			return err
		}
		reqs, limits := node.GetPodsTotalRequestsAndLimits(podList)
		fmt.Printf("\t\t%s\n", n.Name)
		data := make([][]string, 0)
		for k, v := range reqs {
			limit := limits[k]
			allocate := n.Status.Allocatable[k]
			row := make([]string, 0)
			row = append(row, k.String(), v.String(), limit.String(), allocate.String())
			data = append(data, row)
		}
		printTable([]string{"Name", "Requests", "Limits", "Allocatable"}, data)
		fmt.Println()
	}

	return nil
}

func printTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(header)
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.FgGreenColor})

	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.AppendBulk(data)
	table.Render()
}
