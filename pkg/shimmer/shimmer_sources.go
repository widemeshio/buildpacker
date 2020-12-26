package shimmer

import (
	"log"

	"github.com/widemeshcloud/pack-shimmer/pkg/shimmer/sources"
)

// InstallSources installs a list of sources
func (shimmer *Shimmer) InstallSources(sources []sources.Source) {
	log.Printf("installing sources, %d", len(sources))
	shimmer.Sources = append(shimmer.Sources, sources...)
}
