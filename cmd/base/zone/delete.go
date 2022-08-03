package zone

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "zone",
		DisableFlagsInUseLine: true,
		Short:                 "delete zone with the specified zone name from config center",
		Long:                  "",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	return cmd
}
