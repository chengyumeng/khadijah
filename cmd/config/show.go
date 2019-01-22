package config

import (
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/gin-gonic/gin/json"
	"github.com/spf13/cobra"
)

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Used to show all user configurations.",
	Run: func(cmd *cobra.Command, args []string) {
		data, _ := json.MarshalIndent(config.GlobalOption, " ", " ")
		fmt.Println(string(data))
	},
}
