package config

import (
	"github.com/spf13/cobra"
)

// Config show and set command interface
var ConfigCmd = &cobra.Command{
	Use:     "config",
	Short:   "For khadijah config tool.",
	Example: `khadijah config show`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	ConfigCmd.AddCommand(showCmd)
	ConfigCmd.AddCommand(setCmd)

}
