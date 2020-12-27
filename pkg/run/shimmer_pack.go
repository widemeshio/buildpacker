package run

import (
	"context"

	"github.com/widemeshcloud/pack-shimmer/pkg/shimmer/sources"

	"github.com/widemeshcloud/pack-shimmer/pkg/shimmer"
)

// ShimmerPack runs pack with shimmed buildpacks
type ShimmerPack struct {
	Config
}

// Run runs pack command
func (pack *ShimmerPack) Run(ctx context.Context) error {
	shimmer := &shimmer.Shimmer{}
	shimmer.InstallSources(sources.BuiltIn())
	buildpacks, err := shimmer.Apply(ctx, pack.Buildpacks)
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
	return nil
}
