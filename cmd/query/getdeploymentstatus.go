package query

import (
	"github.com/chengyumeng/khadijah/pkg/query"
	"github.com/spf13/cobra"
)

var getDeploymentStatusOption query.GetDeploymentStatusOption

var GetDeploymentStatusCmd = &cobra.Command{
	Use:     "getdeploymentstatus",
	Short:   "Get Deployment Status by OpenAPI",
	Example: "khadijah getpodinfo -c=SHBT --label=app=openapi-demo",
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
