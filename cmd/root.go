package cmd

import (
	"github.com/chengyumeng/khadijah/cmd/config"
	"github.com/chengyumeng/khadijah/cmd/describe"
	"github.com/chengyumeng/khadijah/cmd/exec"
	"github.com/chengyumeng/khadijah/cmd/get"
	"github.com/chengyumeng/khadijah/cmd/loginout"
	"github.com/chengyumeng/khadijah/cmd/query"
	"github.com/spf13/cobra"
)

// 是否输出 version 信息
var v bool

// RootCmd 是命令行的入口
var RootCmd = &cobra.Command{
	Use: "khadijah",
	Long: `The Client Tool for Wayne
Email: 792400644@qq.com 微信公众号: 程天写代码
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if v {
			versionE(cmd, args)
		}
	},
}

func init() {
	cobra.EnableCommandSorting = false
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(config.ConfigCmd)
	RootCmd.AddCommand(loginout.LoginCmd)
	RootCmd.AddCommand(loginout.LogoutCmd)
	RootCmd.AddCommand(get.GetCmd)
	RootCmd.AddCommand(describe.DescribeCmd)
	RootCmd.AddCommand(query.QueryCmd)
	RootCmd.AddCommand(exec.ExecCmd)

	RootCmd.Flags().BoolVarP(&v, "version", "v", false, "show version")
}
