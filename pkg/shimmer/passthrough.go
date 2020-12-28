package shimmer

// Passthrough is a buildpack that reports the cnb pack argument as the same reported buildpack, useful in buildpack references such as "heroku/go" which buildpacker does not apply any shim
type Passthrough struct {
	Buildpack string
}

// PackArgument return argument for passthrough
func (pt *Passthrough) PackArgument() string {
	return pt.Buildpack
}
