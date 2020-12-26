package sources

// Source matches the given buildpack and returns an Unpacker instances
type Source interface {
	Create(buildpack string) Unpacker
}
