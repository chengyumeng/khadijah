package config

import (
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:     "config",
	Short:   "For khadijah config tool.",
	Example: `khadijah config show`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	ConfigCmd.AddCommand(ShowCmd)
	ConfigCmd.AddCommand(SetCmd)

}
