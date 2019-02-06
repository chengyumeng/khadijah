package query

import (
	"github.com/chengyumeng/khadijah/pkg/query"
	"github.com/spf13/cobra"
)

var getVIPInfoOption query.GetVIPInfoOption

var getVIPInfoCmd = &cobra.Command{
	Use:     "getvipinfo",
	Short:   "Get VIP Info by OpenAPI(Only for Qihoo360)",
	Example: "khadijah getvipinfo -p=21273",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := query.NewProxy()
		proxy.GetVIPInfo(getVIPInfoOption)
	},
}

func init() {
	getVIPInfoCmd.Flags().IntVarP(&getVIPInfoOption.Port, "port", "p", 8080, "")
}
