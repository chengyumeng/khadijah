package ssh

import (
	"github.com/spf13/cobra"
)

var ExecCmd = &cobra.Command{
	Use:   "log",
	Short: "ssh in pod",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
