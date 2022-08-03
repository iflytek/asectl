package base

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xfyun/asectl/cmd/base/cluster"
	"github.com/xfyun/asectl/cmd/base/project"
	"github.com/xfyun/asectl/cmd/base/service"
	"github.com/xfyun/asectl/cmd/base/serviceVersion"
	"github.com/xfyun/asectl/cmd/base/zone"
)

var (
	createLong = `
	Create a resource from a file or from stdin.
	`

	createExample = `
	# Create a project with the specified name.
	asectl create project test

	# Create a cluster with the specified name.
	asectl create cluster test
	`

	getLong = `
	Get a resource info from server. 
	`

	getExample = `
	# Get a project info from server
	asectl get project test

	# Get a cluster info from server
	asectl get cluster test
	`

	deleteLong = `
	Delete a resource from server.
	`

	deleteExample = `
	# Delete a project from server
	asectl delete project test

	# Delete a cluster from server
	asectl delete cluster test
	`
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create",
		DisableFlagsInUseLine: true,
		Short:                 "create a resource from a file or stdin.",
		Long:                  createLong,
		Example:               createExample,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	cmd.AddCommand(zone.NewCmdCreate())
	cmd.AddCommand(project.NewCmdCreate())
	cmd.AddCommand(cluster.NewCmdCreate())
	cmd.AddCommand(service.NewCmdCreate())
	cmd.AddCommand(serviceVersion.NewCmdCreate())
	return cmd
}

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "get",
		DisableFlagsInUseLine: true,
		Short:                 "get a resource from server.",
		Long:                  getLong,
		Example:               getExample,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	cmd.AddCommand(zone.NewCmdGet())
	cmd.AddCommand(project.NewCmdGet())
	cmd.AddCommand(cluster.NewCmdGet())
	cmd.AddCommand(service.NewCmdGet())
	cmd.AddCommand(serviceVersion.NewCmdGet())

	return cmd
}

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete",
		DisableFlagsInUseLine: true,
		Short:                 "delete a resource from server.",
		Long:                  deleteLong,
		Example:               deleteExample,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	cmd.AddCommand(zone.NewCmdDelete())
	cmd.AddCommand(project.NewCmdDelete())
	cmd.AddCommand(cluster.NewCmdDelete())
	cmd.AddCommand(service.NewCmdDelete())
	cmd.AddCommand(serviceVersion.NewCmdDelete())

	return cmd
}
