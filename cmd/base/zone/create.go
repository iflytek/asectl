package zone

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "zone",
		DisableFlagsInUseLine: true,
		Short:                 "create a zone with the specified zone name",
		Long:                  "",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	return cmd
}
