package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/widemeshcloud/pack-shimmer/pkg/commands/build"
)

var rootCmd = &cobra.Command{
	Use:   "pack-shimmer",
	Short: "pack-shimmer - Cloud-Native buildpacks + Legacy Shims",
	Long:  "pack-shimmer - builds Cloud-Native buildpacks automatically adding shims for old legacy Heroku buildpacks",
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
