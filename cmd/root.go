package cmd

import (
	"github.com/chengyumeng/khadijah/cmd/describe"
	"github.com/chengyumeng/khadijah/cmd/get"
	"github.com/chengyumeng/khadijah/cmd/log"
	"github.com/chengyumeng/khadijah/cmd/login"
	"github.com/chengyumeng/khadijah/cmd/ssh"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "khadijah",
	Long: `The Client Tool for Wayne
Email: 792400644@qq.com
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(get.GetCmd)
	RootCmd.AddCommand(describe.DescribeCmd)
	RootCmd.AddCommand(log.LogCmd)
	RootCmd.AddCommand(ssh.ExecCmd)
	RootCmd.AddCommand(login.LoginCmd)
}
