package describe

import (
	"github.com/spf13/cobra"
)

var DescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe resource info",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
