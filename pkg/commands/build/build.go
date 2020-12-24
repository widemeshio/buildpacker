package build

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pathArg string
var builderArg string
var buildpacksArg []string

// Command command definition
var Command = &cobra.Command{
	Use:   "build <image-name> [flags]",
	Short: "builds an image",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		imageTag := args[0]
		fmt.Printf("image tag: %s\n", imageTag)
		return fmt.Errorf("h")
	},
}

func init() {
	Command.Flags().StringVarP(&pathArg, "path", "p", ".", "--path <directory>")
	Command.Flags().StringVarP(&builderArg, "builder", "B", "", "Builder image")
	Command.Flags().StringSliceVarP(&buildpacksArg, "buildpack", "b", nil, `
	Buildpack to use. One of:
                                   a buildpack by id and version in the form of '<buildpack>@<version>',
                                   path to a buildpack directory (not supported on Windows),
                                   path/URL to a buildpack .tar or .tgz file, or
                                   a packaged buildpack image name in the form of '<hostname>/<repo>[:<tag>]'
                                 Repeat for each buildpack in order,
								   or supply once by comma-separated list
								   `)
}
