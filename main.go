package main

import (
	"os"

	"github.com/chengyumeng/khadijah/cmd"
	"github.com/chengyumeng/khadijah/pkg/config"
	utilslog "github.com/chengyumeng/khadijah/pkg/utils/log"
)

var (
	Version = ""
)

func main() {
	config.Version = Version
	err := config.LoadOption()
	if err != nil {
		utilslog.AppLogger.Error("Error Load Option")
	}

	err = cmd.RootCmd.Execute()
	if err != nil {
		utilslog.AppLogger.Error("")
		os.Exit(1)
	}
}
