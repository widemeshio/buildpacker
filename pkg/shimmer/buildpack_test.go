package shimmer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildpack(t *testing.T) {
	bps := Buildpacks{
		&testBuildpackArg{
			arg: "heroku/go",
		},
		&testBuildpackArg{
			arg: "heroku/node",
		},
		&testBuildpackArg{
			arg: "heroku/procfile",
		},
	}
	require.Equal(t, []string{"heroku/go", "heroku/node", "heroku/procfile"}, bps.PackArguments())
}

type testBuildpackArg struct {
	arg string
}

func (pt *testBuildpackArg) PackArgument() string {
	return pt.arg
}
