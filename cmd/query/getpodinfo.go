package query

import (
	"github.com/chengyumeng/khadijah/pkg/query"
	"github.com/spf13/cobra"
)

var option query.GetPodInfoOption

var GetPodInfoCmd = &cobra.Command{
	Use:   "getpodinfo",
	Short: "Get Pod Info by OpenAPI",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := query.NewProxy()
		proxy.GetPodInfo(option)
	},
}

func init() {
	GetPodInfoCmd.Flags().StringVarP(&option.Cluster, "cluster", "c", "", "")
	GetPodInfoCmd.Flags().StringVarP(&option.LabelSelector, "label", "l", "", "")
}
