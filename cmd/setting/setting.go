package setting

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "init",
		DisableFlagsInUseLine: true,
		Short:                 "init the asectl",
		Long:                  "",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	return cmd
}
