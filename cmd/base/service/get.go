package service

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "service",
		DisableFlagsInUseLine: true,
		Aliases:               []string{"svc"},
		Short:                 "get a service with the specified service name",
		Long:                  "",
		Example:               "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	return cmd
}
