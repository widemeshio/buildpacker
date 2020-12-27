package shimmer

// Passthrough a passthrough buildpack
type Passthrough struct {
	Buildpack string
}

// PackArgument return argument for passthrough
func (pt *Passthrough) PackArgument() string {
	return pt.Buildpack
}
