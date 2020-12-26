package sources

import "context"

// Unpacker implements unpacking to a local file-system directory of the underlying buildpack
type Unpacker interface {
	Unpack(ctx context.Context, destinationDir string) error
	Buildpack() string
	RequestedVersion() string
}
