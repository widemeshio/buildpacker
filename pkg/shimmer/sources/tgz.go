package sources

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mholt/archiver/v3"
	"github.com/widemeshio/buildpacker/pkg/dl"
)

func init() {
	registerBuiltinSource(tgz)
}

type tgzSource struct{}

var tgz = &tgzSource{}

func (src *tgzSource) Create(buildpack string) Unpacker {
	if strings.HasSuffix(buildpack, ".tgz") || strings.HasSuffix(buildpack, ".tar.gz") {
		return &tgzUnpacker{
			buildpack: buildpack,
		}
	}
	return nil
}

type tgzUnpacker struct {
	buildpack string
}

func (unpacker *tgzUnpacker) CanonicalBuildpack() string {
	return unpacker.buildpack
}

func (unpacker *tgzUnpacker) OriginalBuildpack() string {
	return unpacker.buildpack
}

func (unpacker *tgzUnpacker) RequestedVersion() string {
	return ""
}

func (unpacker *tgzUnpacker) Unpack(ctx context.Context, destinationDir string) error {
	file, err := ioutil.TempFile("", "*.tar.gz")
	if err != nil {
		return fmt.Errorf("failed to create temp file for tgz, %w", err)
	}
	tempFile := file.Name()
	defer os.Remove(tempFile)
	if err := dl.DownloadFile(unpacker.buildpack, tempFile); err != nil {
		return fmt.Errorf("failed to download tgz file, %w", err)
	}
	if err := archiver.Unarchive(tempFile, destinationDir); err != nil {
		return fmt.Errorf("failed to unarchive tgz in destination directory, %w", err)
	}
	return nil
}
