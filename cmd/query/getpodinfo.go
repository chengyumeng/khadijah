package query

import (
	"github.com/chengyumeng/khadijah/pkg/query"
	"github.com/spf13/cobra"
)

var option query.GetPodInfoOption

var getPodInfoCmd = &cobra.Command{
	Use:     "getpodinfo",
	Short:   "Get Pod Info by OpenAPI",
	Example: "khadijah getpodinfo -c=SHBT --label=app=openapi-demo",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := query.NewProxy()
		proxy.GetPodInfo(option)
	},
}

func init() {
	getPodInfoCmd.Flags().StringVarP(&option.Cluster, "cluster", "c", "", "")
	getPodInfoCmd.Flags().StringVarP(&option.LabelSelector, "label", "l", "", "")
}
