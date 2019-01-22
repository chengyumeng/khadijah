package main

import (
	"os"
	"fmt"

	"github.com/chengyumeng/khadijah/cmd"
	"github.com/chengyumeng/khadijah/pkg/config"
	utilslog "github.com/chengyumeng/khadijah/pkg/utils/log"
)

var (
	Version = "beta"
)

func main() {
	err := config.LoadOption()
	if err != nil {
		utilslog.AppLogger.Error("Error Load Option")
	}
	if config.GlobalOption.System == nil {
		err := config.SetSystem(config.System{
			BaseURL: "http://127.0.0.1:4200",
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
