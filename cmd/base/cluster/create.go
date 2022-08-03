package cluster

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "cluster",
		DisableFlagsInUseLine: true,
		Short:                 "create a cluster with the specified cluster name",
		Long:                  "create a cluster with the specified cluster name",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	// TODO: 初始化配置中心管理接口类，并将创建cluster需要的参数绑定到参数上
	var proj string
	cmd.Flags().String("proj", proj, "")
	cmd.MarkFlagRequired("proj")

	return cmd
}
