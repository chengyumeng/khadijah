package main

import (
	"os"

	"github.com/chengyumeng/khadijah/cmd"
	"github.com/chengyumeng/khadijah/pkg/config"
)

var (
	Version = ""
)

func main() {
	config.Version = Version
	err := cmd.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
