package get

import (
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get resource info",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
