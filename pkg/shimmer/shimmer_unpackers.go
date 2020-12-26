package shimmer

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
	unpackers := make([]sources.Unpacker, len(buildpacks))
	for ix, buildpack := range buildpacks {
		unpacker := shimmer.createUnpacker(ctx, buildpack)
		if unpacker == nil {
			return nil, fmt.Errorf("no source was able to unpack buildpack %s", buildpack)
		}
		unpackers[ix] = unpacker
	}
	return unpackers, nil
}

func (shimmer *Shimmer) unpack(ctx context.Context, buildpacks []string) ([]UnpackedBuildpack, error) {
	unpackers, err := shimmer.createUnpackers(ctx, buildpacks)
	if err != nil {
		return nil, fmt.Errorf("failed to build unpackers, %w", err)
	}
	localBuildpacks := make([]UnpackedBuildpack, len(buildpacks))
	for ix, unpacker := range unpackers {
		localBuildpackRoot, err := ioutil.TempDir("", "buildpack-shimmed-*")
		if err != nil {
			return nil, fmt.Errorf("unable to temp dir for buildpack %s, %w", unpacker.Buildpack(), err)
		}
		unpacked := UnpackedBuildpack{
			Unpacker: unpacker,
			LocalDir: localBuildpackRoot,
		}
		targetDir := unpacked.TargetDir()
		if err := os.Mkdir(targetDir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("unable to create target dir %s, %w", unpacker.Buildpack(), err)
		}
		if err := unpacker.Unpack(ctx, targetDir); err != nil {
			return nil, fmt.Errorf("unable to unpack %s, %w", unpacker.Buildpack(), err)
		}
		localBuildpacks[ix] = unpacked
	}
	return localBuildpacks, nil
}

func targetDirOf(buildpackRootDir string) string {
	targetDir := filepath.Join(buildpackRootDir, "target")
	return targetDir
}

// UnpackedBuildpack unpacked buildpack
type UnpackedBuildpack struct {
	Unpacker sources.Unpacker
	LocalDir string
}

// TargetDir returns the target dir of a shimmed buildpack
func (unpacked UnpackedBuildpack) TargetDir() string {
	return targetDirOf(unpacked.LocalDir)
}
