package sources

import (
	"context"
	"fmt"
	"strings"
)

func init() {
	registerBuiltinSource(&herokuSource{})
}

type herokuSource struct{}

func (src *herokuSource) Create(buildpack string) Unpacker {
	if strings.Count(buildpack, "/") != 1 {
		return nil
	}
	parts := strings.SplitN(buildpack, "/", -1)
	registryNamespace := parts[0]
	registryBuildpack := parts[1]
	tgzFile := fmt.Sprintf("https://buildpack-registry.s3.amazonaws.com/buildpacks/%s/%s.tgz", registryNamespace, registryBuildpack)
	return &herokuUnpacker{
		tgz.Create(tgzFile),
		buildpack,
	}
}

type herokuUnpacker struct {
	tgzUnpacker Unpacker
	buildpack   string
}

func (unpacker *herokuUnpacker) Buildpack() string {
	return unpacker.buildpack
}

func (unpacker *herokuUnpacker) RequestedVersion() string {
	return ""
}

func (unpacker *herokuUnpacker) Unpack(ctx context.Context, destinationDir string) error {
	return unpacker.tgzUnpacker.Unpack(ctx, destinationDir)
}
