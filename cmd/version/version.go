package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var goVersion string
var gitCommit string
var branch string
var buildTime string
var oSArch string

func Version() {
	fmt.Printf("Go version:  %v\n", goVersion)
	fmt.Printf("Git commit:  %v\n", gitCommit)
	fmt.Printf("branch:      %v\n", branch)
	fmt.Printf("Built:       %v\n", buildTime)
	fmt.Printf("OS/Arch:     %v\n", oSArch)
}

var (
	versionExample = `
	# Print the client and server versions for the current context
	asectl version
	`
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Short:   "Print the client and server version information",
		Long:    "Print the client and server version information for the current context",
		Example: versionExample,
		Run: func(cmd *cobra.Command, args []string) {
			Version()
		},
	}
	return cmd
}
