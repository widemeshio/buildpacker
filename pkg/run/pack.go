package run

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
	return nil
}
