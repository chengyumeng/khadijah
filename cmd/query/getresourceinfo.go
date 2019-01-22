package query

import (
	"github.com/chengyumeng/khadijah/pkg/query"
	"github.com/spf13/cobra"
)

var getResourceInfoOption query.GetResourceInfoOption

var GetResourceInfoCmd = &cobra.Command{
	Use:   "getresourceinfo",
	Short: "Get Resource Info by OpenAPI",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := query.NewProxy()
		proxy.GetResourceInfo(getResourceInfoOption)
	},
}

func init() {
	GetResourceInfoCmd.Flags().StringVarP(&getResourceInfoOption.Type, "type", "t", "", "")
	GetResourceInfoCmd.Flags().StringVarP(&getResourceInfoOption.Name, "name", "n", "", "")
}
