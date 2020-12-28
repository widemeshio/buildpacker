package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/widemeshcloud/buildpacker/pkg/commands/build"
)

var rootCmd = &cobra.Command{
	Use:   "buildpacker",
	Short: "buildpacker - Cloud-Native buildpacks + Legacy Shims",
	Long:  "buildpacker - builds Cloud-Native buildpacks automatically adding shims for old legacy Heroku buildpacks",
}

func init() {
	rootCmd.AddCommand(build.Command)
}

// Execute executes the main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
