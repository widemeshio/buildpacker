package shimmer

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/widemeshio/buildpacker/pkg/shimmer/sources"
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

func (shimmer *Shimmer) isPassthroughBuildpack(buildpack string) bool {
	return strings.HasPrefix(buildpack, "heroku/")
}

func (shimmer *Shimmer) prepare(ctx context.Context, buildpacks []string) (Buildpacks, error) {
	prepared := make(Buildpacks, len(buildpacks))
	for ix, buildpack := range buildpacks {
		if shimmer.isPassthroughBuildpack(buildpack) {
			prepared[ix] = &Passthrough{
				Buildpack: buildpack,
			}
			continue
		}
		unpacker := shimmer.createUnpacker(ctx, buildpack)
		if unpacker == nil {
			return nil, fmt.Errorf("no source was able to unpack buildpack %s", buildpack)
		}
		localBuildpackRoot, err := ioutil.TempDir("", "buildpack-shimmed-*")
		if err != nil {
			return nil, fmt.Errorf("unable to temp dir for buildpack %s, %w", unpacker.OriginalBuildpack(), err)
		}
		unpacked := UnpackedBuildpack{
			Unpacker: unpacker,
			LocalDir: localBuildpackRoot,
		}
		targetDir := unpacked.TargetDir()
		if err := os.Mkdir(targetDir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("unable to create target dir %s, %w", unpacker.OriginalBuildpack(), err)
		}
		if err := unpacker.Unpack(ctx, targetDir); err != nil {
			return nil, fmt.Errorf("unable to unpack %s, %w", unpacker.OriginalBuildpack(), err)
		}
		if err := markAsExecutables(filepath.Join(targetDir, "bin")); err != nil {
			return nil, fmt.Errorf("unable to mark target buildpack files as executables %s, %w", targetDir, err)
		}
		prepared[ix] = unpacked
	}
	return prepared, nil
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

// PackArgument returns the argument for the pack command
func (unpacked UnpackedBuildpack) PackArgument() string {
	return unpacked.LocalDir
}

func markAsExecutables(dir string) error {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to walk filed to mark as executables, %w", err)
	}
	for _, file := range files {
		if err := os.Chmod(file, 0700); err != nil {
			return fmt.Errorf("failed to change permission of file %s, %w", file, err)
		}
	}
	return nil
}
