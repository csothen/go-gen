package root

import (
	versionCmd "github.com/csothen/go-gen/pkg/cmd/version"
	"github.com/spf13/cobra"
)

func NewCmdRoot(version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-gen <command> [flags]",
		Short: "Code Generation CLI",
		Long:  "Generate code for web applications in various target languages from the command line",
		Example: `
		$ go-gen generate
		$ go-gen version
		$ go-gen validate ./config.json
		`,
	}

	cmd.AddCommand(versionCmd.NewCmdVersion(version, buildDate))

	return cmd
}
