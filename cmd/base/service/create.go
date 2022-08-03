package service

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "service",
		DisableFlagsInUseLine: true,
		Aliases:               []string{"svc"},
		Short:                 "create a service with the specified service name",
		Long:                  "",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	return cmd
}
