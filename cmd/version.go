package cmd

import (
	"fmt"

	"github.com/chengyumeng/khadijah/pkg/config"
	"github.com/spf13/cobra"
	"runtime"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "get version",
	Run:     versionE,
}

var versionData = `khadijah: A client tool for Kubernetes via Wayne.
                                                                   dddddddd
kkkkkkkk          hhhhhhh                                          d::::::d  iiii   jjjj                  hhhhhhh
k::::::k          h:::::h                                          d::::::d i::::i j::::j                 h:::::h
k::::::k          h:::::h                                          d::::::d  iiii   jjjj                  h:::::h
k::::::k          h:::::h                                          d:::::d                                h:::::h
 k:::::k    kkkkkkkh::::h hhhhh         aaaaaaaaaaaaa      ddddddddd:::::d iiiiiiijjjjjjj  aaaaaaaaaaaaa   h::::h hhhhh
 k:::::k   k:::::k h::::hh:::::hhh      a::::::::::::a   dd::::::::::::::d i:::::ij:::::j  a::::::::::::a  h::::hh:::::hhh
 k:::::k  k:::::k  h::::::::::::::hh    aaaaaaaaa:::::a d::::::::::::::::d  i::::i j::::j  aaaaaaaaa:::::a h::::::::::::::hh
 k:::::k k:::::k   h:::::::hhh::::::h            a::::ad:::::::ddddd:::::d  i::::i j::::j           a::::a h:::::::hhh::::::h
 k::::::k:::::k    h::::::h   h::::::h    aaaaaaa:::::ad::::::d    d:::::d  i::::i j::::j    aaaaaaa:::::a h::::::h   h::::::h
 k:::::::::::k     h:::::h     h:::::h  aa::::::::::::ad:::::d     d:::::d  i::::i j::::j  aa::::::::::::a h:::::h     h:::::h
 k:::::::::::k     h:::::h     h:::::h a::::aaaa::::::ad:::::d     d:::::d  i::::i j::::j a::::aaaa::::::a h:::::h     h:::::h
 k::::::k:::::k    h:::::h     h:::::ha::::a    a:::::ad:::::d     d:::::d  i::::i j::::ja::::a    a:::::a h:::::h     h:::::h
k::::::k k:::::k   h:::::h     h:::::ha::::a    a:::::ad::::::ddddd::::::ddi::::::ij::::ja::::a    a:::::a h:::::h     h:::::h
k::::::k  k:::::k  h:::::h     h:::::ha:::::aaaa::::::a d:::::::::::::::::di::::::ij::::ja:::::aaaa::::::a h:::::h     h:::::h
k::::::k   k:::::k h:::::h     h:::::h a::::::::::aa:::a d:::::::::ddd::::di::::::ij::::j a::::::::::aa:::ah:::::h     h:::::h
kkkkkkkk    kkkkkkkhhhhhhh     hhhhhhh  aaaaaaaaaa  aaaa  ddddddddd   dddddiiiiiiiij::::j  aaaaaaaaaa  aaaahhhhhhh     hhhhhhh
                                                                                   j::::j
                                                                         jjjj      j::::j  version   :  %s
                                                                        j::::jj   j:::::j  language  :  %s
                                                                        j::::::jjj::::::j  platform  :  %s/%s
                                                                         jj::::::::::::j   CPU       :  %d
                                                                           jjj::::::jjj    Goroutine :  %d
                                                                              jjjjjj
`

func versionE(cmd *cobra.Command, args []string) {
	fmt.Printf(versionData, config.GlobalOption.System.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), runtime.NumGoroutine())
}
