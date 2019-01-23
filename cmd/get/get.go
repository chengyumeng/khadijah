package get

import (
	"github.com/chengyumeng/khadijah/pkg/get"
	"github.com/spf13/cobra"
)

var option get.Option

var GetCmd = &cobra.Command{
	Use: "get",
	Short: `Display one wayne resource.

Valid resource types include:
* deployment
* daemonset
* statefulset
* service
* cronjob
* namespace
* application

And you can set ns/app as a filter.`,
	Example: `khadijah get --deployment --ns=default`,
	Run: func(cmd *cobra.Command, args []string) {
		proxy := get.NewProxy(option)
		proxy.Get()
	},
}

func init() {
	GetCmd.Flags().StringVarP(&option.NS, "ns", "n", "", "Wayne namespace filter.")
	GetCmd.Flags().StringVarP(&option.App, "app", "a", "", "Wayne application filter.")
	GetCmd.Flags().BoolVarP(&option.Deployment, "deployment", "d", false, "Whether to output deployment list.")
	GetCmd.Flags().BoolVarP(&option.Daemonset, "daemonset", "", false, "Whether to output daemonset list.")
	GetCmd.Flags().BoolVarP(&option.Statefulset, "statefulset", "", false, "Whether to output statefulset list.")
	GetCmd.Flags().BoolVarP(&option.Service, "service", "s", false, "Whether to output service list.")
	GetCmd.Flags().BoolVarP(&option.Cronjob, "cronjob", "c", false, "Whether to output cronjob list.")
	GetCmd.Flags().BoolVarP(&option.Namespace, "namespace", "", false, "Whether to output namespace list.")
	GetCmd.Flags().BoolVarP(&option.Application, "application", "", false, "Whether to output application list.")
}
