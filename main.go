package main

import (
	"os"

	"github.com/chengyumeng/khadijah/cmd"
	"github.com/chengyumeng/khadijah/pkg/config"
	"fmt"
)

var (
	Version = ""
)

func main() {
	config.Version = Version
	opt, err := config.LoadOption()
	if err != nil {
		fmt.Println(err)
	}
	config.GlobalOption = opt

	err = cmd.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
