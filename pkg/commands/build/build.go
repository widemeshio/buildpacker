package build

import (
	"context"
	"fmt"
	"os"

	"github.com/widemeshcloud/pack-shimmer/pkg/run"

	"github.com/spf13/cobra"
)

var pathArg string
var builderArg string
var buildpacksArg []string
var envsArg []string

// Command command definition
var Command = &cobra.Command{
	Use:   "build <image-name> [flags]",
	Short: "builds an image",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(os.Stderr, "buildpacks shimmer building\n")
		imageTag := args[0]
		pack := run.ShimmerPack{}
		pack.Builder = builderArg
		pack.Buildpacks = buildpacksArg
		pack.Env = envsArg
		pack.ImageTag = imageTag
		pack.Path = pathArg
		return pack.Run(context.Background())
	},
}

func init() {
	Command.Flags().StringVarP(&pathArg, "path", "p", ".", "--path <directory>")
	Command.Flags().StringVarP(&builderArg, "builder", "B", "", "Builder image")
	Command.Flags().StringSliceVarP(&envsArg, "env", "e", nil, `
	Build-time environment variable, in the form 'VAR=VALUE' or 'VAR'.
                                 When using latter value-less form, value will be taken from current
                                   environment at the time this command is executed.
                                 This flag may be specified multiple times and will override
                                   individual values defined by --env-file.`)
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