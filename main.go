package main

import (
	"os"

	"github.com/chengyumeng/khadijah/cmd"
	"github.com/chengyumeng/khadijah/pkg/config"
	utilslog "github.com/chengyumeng/khadijah/pkg/utils/log"
)

// khadijah version
var (
	Version = "0.3.0"
)

func main() {
	err := config.LoadOption()
	if err != nil {
		utilslog.AppLogger.Errorln(err)
	}
	if config.GlobalOption.System == nil {
		err := config.SetSystem(config.System{
			BaseURL:      "http://127.0.3.0:4200",
			WebsocketURL: "ws://127.0.3.0:8080",
		})
		if err != nil {
			utilslog.AppLogger.Errorln(err)
		}
	}
	config.GlobalOption.System.Version = Version

	err = cmd.RootCmd.Execute()
	if err != nil {
		utilslog.AppLogger.Error(err)
		os.Exit(1)
	}
}
