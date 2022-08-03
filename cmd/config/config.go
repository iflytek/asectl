package config

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addConfigExample = ``

func NewCmdPush() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "push",
		DisableFlagsInUseLine: true,
		Short:                 "push config to ASE cluster",
		Long:                  "push config to ASE cluster",
		Example:               addConfigExample,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("push config to ASE cluster")
		},
	}
	return cmd
}
