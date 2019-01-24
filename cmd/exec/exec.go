package exec

import (
	"fmt"

	"github.com/spf13/cobra"

	pkgexec "github.com/chengyumeng/khadijah/pkg/exec"
)

var (
	option = pkgexec.Option{}
)

var ExecCmd = &cobra.Command{
	Use:     "exec",
	Short:   "Execute a command in a container.",
	Example: `khadijah exec -p=openapi-demo-9c5bd44b7-xvjpg -c=SHBT --container=php --cmd=whoami -n=default`,
	Args:    cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ssh := pkgexec.NewSocketShell()
		err := ssh.Connect(option)
		if err != nil {
			fmt.Sprintf("创建连接的时候出现异常： %s", err.Error())
			return
		}

		if option.Terminal {
			go ssh.Listen()
			ssh.StdinSend()
		} else {
			ssh.Listen()
		}
	},
}

func init() {
	ExecCmd.Flags().StringVarP(&option.Cluster, "cluster", "c", "", "Wayne cluster name.")
	ExecCmd.Flags().StringVarP(&option.Namespace, "namespace", "n", "default", "Wayne namespace name.")
	ExecCmd.Flags().StringVarP(&option.Pod, "pod", "p", "", "Kubernetes pod name.")
	ExecCmd.Flags().StringVarP(&option.Container, "container", "", "", "Kubernetes container name.")
	ExecCmd.Flags().StringVarP(&option.Cmd, "cmd", "", "/bin/bash", "command")
	ExecCmd.Flags().BoolVarP(&option.Terminal, "terminal", "t", false, "Whether to create a terminal.")
}
