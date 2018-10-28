package log

import (
	"github.com/spf13/cobra"
)

var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "get pod log",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
