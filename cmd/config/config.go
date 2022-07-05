package config

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addConfigExample = ``

func NewCmdAdd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "add config",
		DisableFlagsInUseLine: true,
		Short:                 "add config to ASE cluster",
		Long:                  "add config to ASE cluster",
		Example:               addConfigExample,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("add config to ASE cluster")
		},
	}
	return cmd
}
