package sources

var builtin []Source

func registerBuiltinSource(src Source) {
	builtin = append(builtin, src)
}

// BuiltIn returns a copy of all builtin buildpack sources
func BuiltIn() []Source {
	cp := make([]Source, 0, len(builtin))
	copy(cp, builtin)
	return cp
}
