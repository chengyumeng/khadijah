package query

import (
	"github.com/chengyumeng/khadijah/pkg/query"
	"github.com/spf13/cobra"
)

var getPodFromIPOption query.GetPodFromIPOption

var GetPodInfoFromIPCmd = &cobra.Command{
	Use:   "getpodinfofromip",
	Short: "Get Pod Info From IP by OpenAPI",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := query.NewProxy()
		proxy.GetPodInfoFromIP(getPodFromIPOption)
	},
}

func init() {
	GetPodInfoFromIPCmd.Flags().StringVarP(&getPodFromIPOption.Cluster, "cluster", "c", "", "")
	GetPodInfoFromIPCmd.Flags().StringVarP(&getPodFromIPOption.IP, "ip", "i", "", "")
}
