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
	DescribeCmd.Flags().StringVarP(&option.Daemontset, "daemontset", "", "", "")
	DescribeCmd.Flags().StringVarP(&option.Statefulset, "statefulset", "", "", "")
	DescribeCmd.Flags().StringVarP(&option.Service, "service", "", "", "")
	DescribeCmd.Flags().StringVarP(&option.Configmap, "configmap", "", "", "")
	DescribeCmd.Flags().StringVarP(&option.Ingress, "ingress", "i", "", "")
	DescribeCmd.Flags().StringVarP(&option.Output, "output", "o", "json", fmt.Sprintf("%s,%s,%s", describe.JSON, describe.YAML, describe.PRETTY))
	DescribeCmd.Flags().StringVarP(&option.Cluster, "cluster", "c", "", "")
}
