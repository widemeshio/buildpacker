package shimmer

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/widemeshcloud/pack-shimmer/pkg/shimmer/sources"
)

// Shimmer shims all the specified buildpacks
type Shimmer struct {
	Sources []sources.Source
}

// Apply prepares all the specified buildpacks with a shim and returns path to local directories with shim applied
func (shimmer *Shimmer) Apply(ctx context.Context, buildpacks []string) ([]string, error) {
	log.Printf("sources %v", shimmer.Sources)
	unpackers, err := shimmer.createUnpackers(ctx, buildpacks)
	if err != nil {
		return nil, fmt.Errorf("failed to build unpackers, %w", err)
	}
	localBuildpacks := make([]string, len(buildpacks))
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
