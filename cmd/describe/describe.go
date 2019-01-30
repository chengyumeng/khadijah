package describe

import (
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/describe"
	"github.com/spf13/cobra"
)

var option describe.Option

var DescribeCmd = &cobra.Command{
	Use:     "describe",
	Short:   `Call kubernetes' API via Wayne, and the interface will return the full kubernetes object.`,
	Example: `khadijah describe --deployment=demo --output=pretty`,
	Run: func(cmd *cobra.Command, args []string) {
		proxy := describe.NewProxy(option)
		proxy.Describe()
	},
}

func init() {
	DescribeCmd.Flags().SortFlags = false
	DescribeCmd.Flags().StringVarP(&option.Namespace, "namespace", "n", "", "Wayne namespace(their is some different from kubernetes namespace).")
	DescribeCmd.Flags().StringVarP(&option.Pod, "pod", "", "", "Kubernetes pod name.")
	DescribeCmd.Flags().StringVarP(&option.Deployment, "deployment", "d", "", "Wayne deployment name.")
	DescribeCmd.Flags().StringVarP(&option.Daemontset, "daemontset", "", "", "Wayne daemonset name.")
	DescribeCmd.Flags().StringVarP(&option.Statefulset, "statefulset", "", "", "Wayne statefulset name.")
	DescribeCmd.Flags().StringVarP(&option.Service, "service", "", "", "Wayne service name.")
	DescribeCmd.Flags().StringVarP(&option.Configmap, "configmap", "", "", "Wayne configmap name.")
	DescribeCmd.Flags().StringVarP(&option.Ingress, "ingress", "i", "", "Wayne ingress name.")
	DescribeCmd.Flags().StringVarP(&option.Output, "output", "o", "json", fmt.Sprintf("Support for the following types:%s,%s,%s.", describe.JSON, describe.YAML, describe.PRETTY))
	DescribeCmd.Flags().StringVarP(&option.Cluster, "cluster", "c", "", "Wayne cluster name.")
}
