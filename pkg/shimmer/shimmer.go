package shimmer

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/widemeshcloud/pack-shimmer/pkg/shimmer/sources"
)

// Shimmer shims all the specified buildpacks
type Shimmer struct {
	Sources []sources.Source
}

// Apply prepares all the specified buildpacks with a shim and returns path to local directories with shim applied
func (pack *Shimmer) Apply(ctx context.Context, buildpacks []string) ([]string, error) {
	unpackers, err := pack.createSources(ctx, buildpacks)
	if err != nil {
		return nil, fmt.Errorf("failed to build unpackers")
	}
	localBuildpacks := make([]string, 0, len(buildpacks))
	for ix, unpacker := range unpackers {
		destinationDir, err := ioutil.TempDir("", "buildpack-shimmed-*")
		if err != nil {
			return nil, fmt.Errorf("unable to temp dir for buildpack %s, %w", unpacker.Buildpack(), err)
		}
		if err := unpacker.Unpack(ctx, destinationDir); err != nil {
			return nil, fmt.Errorf("unable to unpack %s, %w", unpacker.Buildpack(), err)
		}
		localBuildpacks[ix] = destinationDir
	}
	return localBuildpacks, nil
}

func (pack *Shimmer) createSource(ctx context.Context, buildpack string) sources.Unpacker {
	for _, src := range pack.Sources {
		unpacker := src.Create(buildpack)
		if unpacker != nil {
			return unpacker
		}
	}
	return nil
}

func (pack *Shimmer) createSources(ctx context.Context, buildpacks []string) ([]sources.Unpacker, error) {
	unpackers := make([]sources.Unpacker, 0, len(buildpacks))
	for ix, buildpack := range buildpacks {
		unpacker := pack.createSource(ctx, buildpack)
		if unpacker == nil {
			return nil, fmt.Errorf("no source was able to unpack buildpack %s", buildpack)
		}
		unpackers[ix] = unpacker
	}
	return unpackers, nil
}
