package describe

import (
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/describe"
	"github.com/spf13/cobra"
)

var option describe.Option

var DescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe resource info",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := describe.NewProxy(option)
		proxy.Describe()
	},
}

func init() {
	DescribeCmd.Flags().StringVarP(&option.Namespace, "namespace", "n", "", "")
	DescribeCmd.Flags().StringVarP(&option.Deployment, "deployment", "d", "", "")
	DescribeCmd.Flags().StringVarP(&option.Output, "output", "o", "json", fmt.Sprintf(describe.JSON, describe.YAML, describe.PRETTY))
	DescribeCmd.Flags().StringVarP(&option.Cluster, "cluster", "c", "", "")
}
