package serviceVersion

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "service-version",
		DisableFlagsInUseLine: true,
		Short:                 "delete a service-version with the specified version number",
		Long:                  "",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	return cmd
}
