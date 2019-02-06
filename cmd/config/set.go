package config

import (
	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/chengyumeng/khadijah/pkg/utils/log"

	"github.com/spf13/cobra"
)

var (
	apikey string
	system config.System
	setCmd = &cobra.Command{
		Use:     "set",
		Short:   "Used to set some user configurations,such as `APIKey`.",
		Example: `khadijah config set --apikey=example --websocketurl=ws://127.0.0.1:8080 --baseurl=http://127.0.0.1:4200`,
		Run: func(cmd *cobra.Command, args []string) {
			e := false
			if len(apikey) > 0 {
				e = true
				config.SetAPIKey(apikey)
			}
			if len(system.BaseURL) > 0 {
				e = true
				config.GlobalOption.System.BaseURL = system.BaseURL
			}
			if len(system.WebsocketURL) > 0 {
				e = true
				config.GlobalOption.System.WebsocketURL = system.WebsocketURL
			}
			if err := config.Save(); err != nil {
				log.CmdLogger.Errorln(err)
			} else if e {
				log.CmdLogger.Infoln("Success Save Config!")
			} else {
				log.CmdLogger.Warningln("Don't Save Anything!")
			}
		},
	}
)

func init() {
	setCmd.Flags().StringVarP(&apikey, "apikey", "a", "", "In Wayne, APIKEY is primarily used for calls to OpenAPI.")
	setCmd.Flags().StringVarP(&system.WebsocketURL, "websocketurl", "w", "", "Specify Wayne's websocket API call url.")
	setCmd.Flags().StringVarP(&system.BaseURL, "baseurl", "b", "", "Specify Wayne's generic API call url.")
}
