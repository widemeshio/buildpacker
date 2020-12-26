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

const gitURLVersionMatch = ".git#"

type gitSource struct{}

func (src *gitSource) Create(buildpack string) Unpacker {
	if strings.Contains(buildpack, gitURLVersionMatch) {
		index := strings.Index(buildpack, gitURLVersionMatch)
		version := buildpack[index+len(gitURLVersionMatch):]
		return &gitUnpacker{
			buildpack: buildpack[:index],
			version:   version,
		}
	}
	if strings.HasSuffix(buildpack, ".git") {
		return &gitUnpacker{
			buildpack: buildpack,
		}
	}
	return nil
}

type gitUnpacker struct {
	buildpack string
	version   string
}

func (unpacker *gitUnpacker) Buildpack() string {
	return unpacker.buildpack
}

func (unpacker *gitUnpacker) RequestedVersion() string {
	return unpacker.version
}

func (unpacker *gitUnpacker) Unpack(ctx context.Context, destinationDir string) error {
	cmd := exec.Command("git", "clone", unpacker.buildpack, destinationDir)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	if v := unpacker.RequestedVersion(); v != "" {
		cmd := exec.Command("git", "-c", "advice.detachedHead=false", "checkout", v)
		cmd.Dir = destinationDir
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
