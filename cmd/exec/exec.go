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
	Short:   "执行容器命令操作",
	Example: `khadijah exec -it a b c /bin/bash`,
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
	ExecCmd.Flags().StringVarP(&option.Cluster, "cluster", "c", "", "服务所在集群")
	ExecCmd.Flags().StringVarP(&option.Namespace, "namespace", "n", "default", "服务所在命名空间")
	ExecCmd.Flags().StringVarP(&option.Deployment, "deployment", "d", "", "服务的部署名称")
	ExecCmd.Flags().StringVarP(&option.Pod, "pod", "p", "", "服务的 pod 名称")
	ExecCmd.Flags().StringVarP(&option.Container, "container", "", "", "服务中的特定容器名称")
	ExecCmd.Flags().StringVarP(&option.Cmd, "cmd", "", "/bin/bash", "进入服务执行的命令")
	ExecCmd.Flags().BoolVarP(&option.Terminal, "terminal", "t", false, "")
}
