package main

import (
	"fmt"
	"os"

	"github.com/chengyumeng/khadijah/cmd"
	"github.com/chengyumeng/khadijah/pkg/config"
	utilslog "github.com/chengyumeng/khadijah/pkg/utils/log"
)

var (
	Version = "0.0.2"
)

func main() {
	err := config.LoadOption()
	if err != nil {
		utilslog.AppLogger.Error("Error Load Option")
	}
	if config.GlobalOption.System == nil {
		err := config.SetSystem(config.System{
			BaseURL:      "http://127.0.0.2:4200",
			WebsocketURL: "ws://127.0.0.2:8080",
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	config.GlobalOption.System.Version = Version

	err = cmd.RootCmd.Execute()
	if err != nil {
		utilslog.AppLogger.Error("")
		os.Exit(1)
	}
}
