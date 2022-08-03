package cluster

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "cluster",
		DisableFlagsInUseLine: true,
		Short:                 "delete a cluster with the specified cluster name",
		Long:                  "",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	return cmd
}
