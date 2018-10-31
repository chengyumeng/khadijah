package exec

import (
	"github.com/chengyumeng/khadijah/pkg/exec"
	"github.com/spf13/cobra"
)

var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "ssh in pod",
	Run: func(cmd *cobra.Command, args []string) {
		exec.Test()
	},
}
