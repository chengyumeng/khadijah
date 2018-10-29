package get

import (
	"github.com/spf13/cobra"
	"github.com/chengyumeng/khadijah/pkg/get"
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
	GetCmd.Flags().StringVarP(&option.Resource, "resource", "r", "", "Resource Type")
	GetCmd.Flags().StringVarP(&option.Namespace, "namespace", "n", "", "")
	GetCmd.Flags().StringVarP(&option.App, "app", "a", "", "")
}
