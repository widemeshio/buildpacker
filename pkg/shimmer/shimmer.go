package shimmer

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/widemeshcloud/pack-shimmer/pkg/shimmer/sources"
)

// DefaultBuildpackAPIVersion the default API version to write to buildpack.toml
const DefaultBuildpackAPIVersion = "0.4"

// DefaultBuildpackStacks the stacks to write to buildpack.toml
func DefaultBuildpackStacks() []string {
	return []string{"heroku-18", "heroku-20"}
}

// Shimmer shims all the specified buildpacks
type Shimmer struct {
	Sources    []sources.Source
	APIVersion string
	Stacks     []string
}

// BuildpackAPIVersion the BuildpackAPIVersion to use in buildpack.toml
func (shimmer *Shimmer) BuildpackAPIVersion() string {
	if v := shimmer.APIVersion; v != "" {
		return v
	}
	return DefaultBuildpackAPIVersion
}

// BuildpackStacks the stacks to use in buildpack.toml
func (shimmer *Shimmer) BuildpackStacks() []string {
	if v := shimmer.Stacks; len(v) > 0 {
		return shimmer.Stacks
	}
	return DefaultBuildpackStacks()
}

// Apply prepares all the specified buildpacks with a shim and returns path to local directories with shim applied
func (shimmer *Shimmer) Apply(ctx context.Context, buildpacks []string) (ShimmedBuildpacks, error) {
	localBuildpacks, err := shimmer.unpack(ctx, buildpacks)
	if err != nil {
		return nil, err
	}
	shimmedBuildpacks := make(ShimmedBuildpacks, len(localBuildpacks))
	for ix, unpacked := range localBuildpacks {
		// targetDir := unpacked.TargetDir()
		shimmedBuildpack := ShimmedBuildpack{
			UnpackedBuildpack: unpacked,
		}
		tomlFile := shimmedBuildpack.ShimBuildpackToml()
		tomlContent := &bytes.Buffer{}
		err := buildpackTomlTemplate.Execute(tomlContent, &buildpackTomlTemplateParams{
			APIID:   shimmer.BuildpackAPIVersion(),
			ID:      shimmedBuildpack.Unpacker.Buildpack(),
			Name:    shimmedBuildpack.Unpacker.Buildpack(),
			Version: "0.1",
			Stacks:  shimmer.BuildpackStacks(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create buildpack.toml content, %w", err)
		}
		if err := ioutil.WriteFile(tomlFile, tomlContent.Bytes(), os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create buildpack.toml, %w", err)
		}
		shimmedBuildpacks[ix] = shimmedBuildpack
	}
	return shimmedBuildpacks, nil
}

var buildpackToml = `
api = "{{.APIID}}"

[buildpack]
id = "{{.ID}}"
version = "{{.Version}}"
name = "{{.Name}}"

{{range .Stacks}}
[[stacks]]
id = "{{.}}"
{{end}}
`

var buildpackTomlTemplate = template.Must(template.New("toml").Parse(buildpackToml))

type buildpackTomlTemplateParams struct {
	APIID   string
	ID      string
	Name    string
	Version string
	Stacks  []string
}

// ShimmedBuildpack holds information about a shimmed buildpack
type ShimmedBuildpack struct {
	UnpackedBuildpack
}

// ShimBuildpackToml returns the path to the buildpack.toml of the shim
func (shimmed ShimmedBuildpack) ShimBuildpackToml() string {
	return filepath.Join(shimmed.LocalDir, "buildpack.toml")
}

// ShimmedBuildpacks slice of buildpacks with shim applied
type ShimmedBuildpacks []ShimmedBuildpack

// LocalDirs returns the names of the local buildpacks
func (buildpacks ShimmedBuildpacks) LocalDirs() []string {
	dirs := make([]string, len(buildpacks))
	for ix, bp := range buildpacks {
		dirs[ix] = bp.LocalDir
	}
	return dirs
}
