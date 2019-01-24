package config

import (
	"github.com/chengyumeng/khadijah/pkg/config"

	"github.com/spf13/cobra"
)

var APIkey string

var System config.System

var SetCmd = &cobra.Command{
	Use:     "set",
	Short:   "Used to set some user configurations,such as `APIKey`.",
	Example: `khadijah config set --apikey=example --websocketurl=ws://127.0.0.1:8080 --baseurl=http://127.0.0.1:4200`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(APIkey) > 0 {
			config.SetAPIKey(APIkey)
		}
		if len(System.BaseURL) > 0 {
			config.GlobalOption.System.BaseURL = System.BaseURL
		}

		if len(System.WebsocketURL) > 0 {
			config.GlobalOption.System.WebsocketURL = System.WebsocketURL
		}
		config.Save()
	},
}

func init() {
	SetCmd.Flags().StringVarP(&APIkey, "apikey", "a", "", "In Wayne, APIKEY is primarily used for calls to OpenAPI.")
	SetCmd.Flags().StringVarP(&System.WebsocketURL, "websocketurl", "w", "", "Specify Wayne's websocket API call url.")
	SetCmd.Flags().StringVarP(&System.BaseURL, "baseurl", "b", "", "Specify Wayne's generic API call url.")
}
