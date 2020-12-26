package sources

import (
	"context"
	"os"
	"os/exec"
	"strings"
)

func init() {
	registerBuiltinSource(&gitSource{})
}

type gitSource struct{}

func (src *gitSource) Create(buildpack string) Unpacker {
	if strings.HasSuffix(buildpack, ".git") {
		return &gitUnpacker{buildpack}
	}
	return nil
}

type gitUnpacker struct {
	buildpack string
}

func (unpacker *gitUnpacker) Buildpack() string {
	return unpacker.buildpack
}

func (unpacker *gitUnpacker) Unpack(ctx context.Context, destinationDir string) error {
	cmd := exec.Command("git", "clone", unpacker.buildpack, destinationDir)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
