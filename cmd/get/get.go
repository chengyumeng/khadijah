package get

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/chengyumeng/khadijah/pkg/get"
	"github.com/chengyumeng/khadijah/pkg/utils/resource"
)

var option get.Option

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: `Display one wayne resource.`,
	Long: `Display one wayne resource.

Valid resource types include:
* deployment
* daemonset
* statefulset
* service
* cronjob
* namespace
* application
* apikey

And you can set ns/app as a filter.`,
	Example: `khadijah get deployment --ns=default`,
	Run: func(cmd *cobra.Command, args []string) {
		option.Option = resource.ParserArgs(args)
		proxy := get.NewProxy(option)
		proxy.Get()
	},
}

func init() {
	GetCmd.Flags().SortFlags = false
	GetCmd.Flags().StringVarP(&option.NS, "namespace", "n", "", "Wayne namespace filter.")
	GetCmd.Flags().StringVarP(&option.App, "application", "a", "", "Wayne application filter.")
	GetCmd.Flags().StringVarP(&option.Output, "output", "o", "pretty", fmt.Sprintf("Support for the following types:%s,%s.", get.ROW, get.PRETTY))
}
