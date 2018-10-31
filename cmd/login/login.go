package login

import (
	"github.com/chengyumeng/khadijah/pkg/login"
	"github.com/spf13/cobra"
)

var option login.Option

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to the wayne platform for more resource access.",
	Run: func(cmd *cobra.Command, args []string) {
		login.Login(option)
	},
}

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from the wayne platform.",
	Run: func(cmd *cobra.Command, args []string) {
		login.Clear()
	},
}

func init() {
	LoginCmd.Flags().StringVarP(&option.Username, "username", "u", "", "User Name")
	LoginCmd.Flags().StringVarP(&option.Password, "password", "p", "", "User Password")

}
