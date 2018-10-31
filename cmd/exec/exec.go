package exec

import (
	"github.com/spf13/cobra"
	"github.com/chengyumeng/khadijah/pkg/exec"
)

var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "ssh in pod",
	Run: func(cmd *cobra.Command, args []string) {
		exec.Test()
	},
}
