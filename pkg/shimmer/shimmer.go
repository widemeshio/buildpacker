package shimmer

import (
	"context"

	"github.com/widemeshcloud/pack-shimmer/pkg/shimmer/sources"
)

// Shimmer shims all the specified buildpacks
type Shimmer struct {
	Sources []sources.Source
}

// Apply prepares all the specified buildpacks with a shim and returns path to local directories with shim applied
func (shimmer *Shimmer) Apply(ctx context.Context, buildpacks []string) ([]string, error) {
	localBuildpacks, err := shimmer.unpack(ctx, buildpacks)
	if err != nil {
		return nil, err
	}
	return localBuildpacks, nil
}
