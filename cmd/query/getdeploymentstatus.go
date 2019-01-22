package query

import (
	"github.com/chengyumeng/khadijah/pkg/query"
	"github.com/spf13/cobra"
)

var getDeploymentStatusOption query.GetDeploymentStatusOption

var GetDeploymentStatusCmd = &cobra.Command{
	Use:   "getdeploymentstatus",
	Short: "Get Resource Info by OpenAPI",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := query.NewProxy()
		proxy.GetDeploymentStatus(getDeploymentStatusOption)
	},
}

func init() {
	GetDeploymentStatusCmd.Flags().StringVarP(&getDeploymentStatusOption.Deployment, "deployment", "d", "", "")
	GetDeploymentStatusCmd.Flags().StringVarP(&getDeploymentStatusOption.Namespace, "namespace", "n", "", "")
	GetDeploymentStatusCmd.Flags().StringVarP(&getDeploymentStatusOption.Cluster, "cluster", "c", "", "")
}
