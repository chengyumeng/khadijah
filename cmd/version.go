package cmd

import (
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/spf13/cobra"
	"runtime"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "get version",
	Run:     versionE,
}

var versionData = `khadijah: A client tool for Kubernetes via Wayne.
version:  %s
language: %s
os info:  %s/%s`

func versionE(cmd *cobra.Command, args []string) {
	fmt.Printf(versionData, config.GlobalOption.System.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
