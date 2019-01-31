package query

import (
	"github.com/chengyumeng/khadijah/pkg/utils/log"
	"github.com/spf13/cobra"
)

var QueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query by Wayne OpenAPI",
	Run: func(cmd *cobra.Command, args []string) {
		log.CmdLogger.Infoln("You should insert a correct child command!")
	},
}

func init() {
	QueryCmd.AddCommand(GetVIPInfoCmd)
	QueryCmd.AddCommand(GetPodInfoCmd)
	QueryCmd.AddCommand(GetPodInfoFromIPCmd)
	QueryCmd.AddCommand(GetResourceInfoCmd)
	QueryCmd.AddCommand(GetDeploymentStatusCmd)
}
