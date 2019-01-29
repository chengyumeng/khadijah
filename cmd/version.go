package cmd

import (
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "get version",
	Run:     versionE,
}

func versionE(cmd *cobra.Command, args []string) {
	fmt.Printf("khadijah %s \n", config.GlobalOption.System.Version)
}
