package get

import (
	"github.com/chengyumeng/khadijah/pkg/get"
	"github.com/spf13/cobra"
)

var option get.Option

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get resource info",
	Run: func(cmd *cobra.Command, args []string) {
		proxy := get.NewProxy(option)
		proxy.Get()
	},
}

func init() {
	GetCmd.Flags().StringVarP(&option.NS, "ns", "n", "", "")
	GetCmd.Flags().StringVarP(&option.App, "app", "a", "", "")
	GetCmd.Flags().BoolVarP(&option.Deployment, "deployment", "d", false, "")
	GetCmd.Flags().BoolVarP(&option.Daemonset, "daemonset", "", false, "")
	GetCmd.Flags().BoolVarP(&option.Statefulset, "statefulset", "", false, "")
	GetCmd.Flags().BoolVarP(&option.Service, "service", "s", false, "")
	GetCmd.Flags().BoolVarP(&option.Cronjob, "cronjob", "c", false, "")
	GetCmd.Flags().BoolVarP(&option.Namespace, "namespace", "", false, "")
	GetCmd.Flags().BoolVarP(&option.Application, "application", "", false, "")
}
