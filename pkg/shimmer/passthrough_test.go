package shimmer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassthrough(t *testing.T) {
	bp := Passthrough{
		Buildpack: "heroku/go",
	}
	require.Equal(t, "heroku/go", bp.PackArgument())
}
