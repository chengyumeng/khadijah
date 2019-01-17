package describe

import (
	"github.com/spf13/cobra"
	"github.com/chengyumeng/khadijah/pkg/describe"
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
	DescribeCmd.Flags().StringVarP(&option.Resource, "resource", "r", "", "Resource Type")
	DescribeCmd.Flags().StringVarP(&option.Namespace, "namespace", "n", "", "")
	DescribeCmd.Flags().StringVarP(&option.App, "app", "a", "", "")
}
