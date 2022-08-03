package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xfyun/asectl/cmd/base"
	"github.com/xfyun/asectl/cmd/config"
	"github.com/xfyun/asectl/cmd/schema"
	"github.com/xfyun/asectl/cmd/setting"
	"github.com/xfyun/asectl/cmd/templates"
	"github.com/xfyun/asectl/cmd/version"
)

var describe = `asectl controls the ASE cluster manager

Find more information at:
    https://github.com/xfyun/asectl
`

var rootCommand = &cobra.Command{
	Use:   "asectl",
	Short: "asectl controls the ASE cluster manager",
	Long:  describe,
	Run:   runHelp,
}

func NewDefaultCommand() *cobra.Command {
	groups := templates.CommandGroups{
		{
			Message: "Init Setting Commands",
			Commands: []*cobra.Command{
				setting.NewCmdInit(),
			},
		},
		{
			Message: "Basic Commands",
			Commands: []*cobra.Command{
				base.NewCmdCreate(),
				base.NewCmdGet(),
				base.NewCmdDelete(),
			},
		},
		{
			Message: "Config Commands",
			Commands: []*cobra.Command{
				config.NewCmdPush(),
			},
		},
		{
			Message: "Schema Commands",
			Commands: []*cobra.Command{
				schema.NewCmdMake(),
			},
		},
	}
	groups.Add(rootCommand)

	filters := []string{"options"}
	templates.ActsAsRootCommand(rootCommand, filters, groups...)

	rootCommand.AddCommand(version.NewCmdVersion())

	return rootCommand
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
