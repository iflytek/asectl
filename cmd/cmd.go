package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xfyun/asectl/cmd/config"
	"github.com/xfyun/asectl/cmd/templates"
	"github.com/xfyun/asectl/cmd/version"
)

var describe = `asectl controls the ASE cluster manager

Find more information at:
    https://xxxxxx
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
			Message: "Config Command",
			Commands: []*cobra.Command{
				config.NewCmdAdd(),
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
