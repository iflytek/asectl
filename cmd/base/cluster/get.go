package cluster

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "cluster",
		DisableFlagsInUseLine: true,
		Short:                 "get a cluster with the specified cluster name",
		Long:                  "",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	return cmd
}
