package query

import (
	"github.com/chengyumeng/khadijah/pkg/query"
	"github.com/spf13/cobra"
)

var getResourceInfoOption query.GetResourceInfoOption

var getResourceInfoCmd = &cobra.Command{
	Use:     "getresourceinfo",
	Short:   `Get Resource Info by OpenAPI`,
	Example: "khadijah getresourceinfo -n=demo --type=deployment",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := query.NewProxy()
		proxy.GetResourceInfo(getResourceInfoOption)
	},
}

func init() {
	getResourceInfoCmd.Flags().StringVarP(&getResourceInfoOption.Type, "type", "t", "", "kubernetes resource type:deployment,daemonset,statefulset etc.")
	getResourceInfoCmd.Flags().StringVarP(&getResourceInfoOption.Name, "name", "n", "", "kubernetes resource name")
}
