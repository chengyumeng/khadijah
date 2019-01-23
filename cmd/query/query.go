package query

import (
	"github.com/spf13/cobra"
)

var QueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query by Wayne OpenAPI",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	QueryCmd.AddCommand(GetPodInfoCmd)
	QueryCmd.AddCommand(GetPodInfoFromIPCmd)
	QueryCmd.AddCommand(GetResourceInfoCmd)
	QueryCmd.AddCommand(GetDeploymentStatusCmd)
}
