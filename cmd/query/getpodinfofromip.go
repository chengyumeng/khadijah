package query

import (
	"github.com/chengyumeng/khadijah/pkg/query"
	"github.com/spf13/cobra"
)

var getPodFromIPOption query.GetPodFromIPOption

var getPodInfoFromIPCmd = &cobra.Command{
	Use:     "getpodinfofromip",
	Short:   "Get Pod Info From IP by OpenAPI",
	Example: "khadijah getpodinfofromip -c=SHBT --ip=172.17.240.119",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := query.NewProxy()
		proxy.GetPodInfoFromIP(getPodFromIPOption)
	},
}

func init() {
	getPodInfoFromIPCmd.Flags().StringVarP(&getPodFromIPOption.Cluster, "cluster", "c", "", "")
	getPodInfoFromIPCmd.Flags().StringVarP(&getPodFromIPOption.IP, "ip", "i", "", "pod ip list")
}
