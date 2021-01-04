package shimmer

// IDDictionary buildpack ID dictionary, each key is the original buildpack ID and the value is what it was replaced for during the shim process
type IDDictionary map[string]string
