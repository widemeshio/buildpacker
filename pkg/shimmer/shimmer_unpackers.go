package shimmer

import (
	"context"
	"fmt"

	"github.com/widemeshcloud/pack-shimmer/pkg/shimmer/sources"
)

func (shimmer *Shimmer) createUnpacker(ctx context.Context, buildpack string) sources.Unpacker {
	for _, src := range shimmer.Sources {
		unpacker := src.Create(buildpack)
		if unpacker != nil {
			return unpacker
		}
	}
	return nil
}

func (shimmer *Shimmer) createUnpackers(ctx context.Context, buildpacks []string) ([]sources.Unpacker, error) {
	unpackers := make([]sources.Unpacker, 0, len(buildpacks))
	for ix, buildpack := range buildpacks {
		unpacker := shimmer.createUnpacker(ctx, buildpack)
		if unpacker == nil {
			return nil, fmt.Errorf("no source was able to unpack buildpack %s", buildpack)
		}
		unpackers[ix] = unpacker
	}
	return unpackers, nil
}
