package run

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/widemeshio/buildpacker/pkg/shimmer/sources"

	"github.com/widemeshio/buildpacker/pkg/shimmer"
)

// ShimmerPack runs pack with shimmed buildpacks
type ShimmerPack struct {
	Config
	IDFile string // JSON file to write the ID of each buildpack
}

// Run runs pack command
func (pack *ShimmerPack) Run(ctx context.Context) error {
	shimmer := &shimmer.Shimmer{}
	shimmer.InstallSources(sources.BuiltIn())
	buildpacks, ids, err := shimmer.Apply(ctx, pack.Buildpacks)
	if err != nil {
		return err
	}
	config := pack.Config
	config.Buildpacks = buildpacks.PackArguments()
	runner := &CnbPack{
		Config: config,
	}
	if err := runner.Run(ctx); err != nil {
		return err
	}
	if filename := pack.IDFile; filename != "" {
		idsContent, err := json.Marshal(ids)
		if err != nil {
			return fmt.Errorf("failed to marshal ID content, %w", err)
		}
		if err := ioutil.WriteFile(filename, idsContent, os.ModePerm); err != nil {
			return fmt.Errorf("failed to write ID file content, %w", err)
		}
	}
	return nil
}
