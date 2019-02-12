package query

import (
	"github.com/chengyumeng/khadijah/pkg/utils/log"
	"github.com/spf13/cobra"
)

// QueryCmd is the interface of query wayne openapi
var QueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query by Wayne OpenAPI",
	Run: func(cmd *cobra.Command, args []string) {
		log.CmdLogger.Infoln("You should insert a correct child command!")
	},
}

func init() {
	QueryCmd.AddCommand(getVIPInfoCmd)
	QueryCmd.AddCommand(getPodInfoCmd)
	QueryCmd.AddCommand(getPodInfoFromIPCmd)
	QueryCmd.AddCommand(getResourceInfoCmd)
	QueryCmd.AddCommand(getDeploymentStatusCmd)
}
