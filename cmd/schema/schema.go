package schema

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdMake() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "make schema from a python wrapper file",
		DisableFlagsInUseLine: true,
		Short:                 "make schema from a python wrapper file",
		Long:                  "make schema from a python wrapper file",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("make schema from a python wrapper file")
		},
	}
	return cmd
}
