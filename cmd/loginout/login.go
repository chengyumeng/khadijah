package loginout

import (
	"github.com/chengyumeng/khadijah/pkg/login"
	"github.com/chengyumeng/khadijah/pkg/utils/log"

	"github.com/spf13/cobra"
)

var (
	option login.Option
    // LoginCmd is the interface of User login command.
	LoginCmd = &cobra.Command{
		Use:     "login",
		Short:   "Log in to the Wayne platform for more API access.",
		Example: "khadijah login -uadmin -padmin",
		Run: func(cmd *cobra.Command, args []string) {
			err := login.Login(option)
			if err != nil {
				log.CmdLogger.Errorln(err)
			}
		},
	}
	// LogoutCmd is the interface of User logout command.
	LogoutCmd = &cobra.Command{
		Use:     "logout",
		Short:   "Log out from the Wayne platform.",
		Example: "khadijah logout",
		Run: func(cmd *cobra.Command, args []string) {
			login.Clear()
		},
	}
)

func init() {
	LoginCmd.Flags().StringVarP(&option.Username, "username", "u", "", "User Name")
	LoginCmd.Flags().StringVarP(&option.Password, "password", "p", "", "User Password")

}
