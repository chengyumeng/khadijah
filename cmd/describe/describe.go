package describe

import (
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/describe"
	"github.com/chengyumeng/khadijah/pkg/utils/resource"
	"github.com/spf13/cobra"
)

var option describe.Option

var DescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: `Call kubernetes' API via Wayne, and the interface will return the full kubernetes object.`,
	Long: `Call kubernetes' API via Wayne, and the interface will return the full kubernetes object.

Valid kubernetes resource obejct types include:
* pod
* deployment
* daemonset
* statefulset
* service
* configmap
* ingress

And you can set ns/app as a filter.`,
	Example: `khadijah describe --deployment=demo --output=pretty`,
	Run: func(cmd *cobra.Command, args []string) {
		option.Option = resource.ParserResource(args)
		proxy := describe.NewProxy(option)
		proxy.Describe()
	},
}

func init() {
	DescribeCmd.Flags().SortFlags = false
	DescribeCmd.Flags().StringVarP(&option.Namespace, "namespace", "n", "", "Wayne namespace(their is some different from kubernetes namespace).")
	DescribeCmd.Flags().StringVarP(&option.Output, "output", "o", "json", fmt.Sprintf("Support for the following types:%s,%s,%s,%s.", describe.JSON, describe.YAML, describe.PRETTY, describe.ROW))
	DescribeCmd.Flags().StringVarP(&option.Cluster, "cluster", "c", "", "Wayne cluster name.")
}
