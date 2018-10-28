package login

import (
	"github.com/spf13/cobra"
	"github.com/chengyumeng/khadijah/pkg/login"
)

var option login.Option

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "ssh in pod",
	Run: func(cmd *cobra.Command, args []string) {
		login.Login(option)
	},
}

func init() {
	LoginCmd.Flags().StringVarP(&option.Username, "username", "u", "", "User Name")
	LoginCmd.Flags().StringVarP(&option.Password, "password", "p", "", "User Password")

}
