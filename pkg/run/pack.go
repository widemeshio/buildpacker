package run

import (
	"os"
	"os/exec"
)

// Config runs the pack command
type Config struct {
	Path       string
	Builder    string
	ImageTag   string
	Buildpacks []string
	Env        []string
}

// Pack runs pack
type Pack struct {
	Config
}

// Run runs pack command
func (pack *Pack) Run() error {
	args := []string{"build", "--path", pack.Path, "--trust-builder", "--builder", pack.Builder}
	for _, bp := range pack.Buildpacks {
		args = append(args, "--buildpack", bp)
	}
	args = append(args, pack.ImageTag)
	cmd := exec.Command("pack", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
