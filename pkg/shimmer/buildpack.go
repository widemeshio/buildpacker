package shimmer

// Buildpack implements PackArgument for the underlying buildpack
type Buildpack interface {
	PackArgument() string
}

// Buildpacks slice of buildpacks with shim applied
type Buildpacks []Buildpack

// PackArguments returns the arguments for the pack command
func (buildpacks Buildpacks) PackArguments() []string {
	dirs := make([]string, len(buildpacks))
	for ix, bp := range buildpacks {
		dirs[ix] = bp.PackArgument()
	}
	return dirs
}
