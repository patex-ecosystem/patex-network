package pt_service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
)

func TestCLIFlagsToEnvVars(t *testing.T) {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:   "test",
			EnvVar: "PT_NODE_TEST_VAR",
		},
		cli.IntFlag{
			Name: "no env var",
		},
	}
	res := cliFlagsToEnvVars(flags)
	require.Contains(t, res, "PT_NODE_TEST_VAR")
}

func TestValidateEnvVars(t *testing.T) {
	provided := []string{"PT_BATCHER_CONFIG=true", "PT_BATCHER_FAKE=false", "LD_PRELOAD=/lib/fake.so"}
	defined := map[string]struct{}{
		"PT_BATCHER_CONFIG": {},
		"PT_BATCHER_OTHER":  {},
	}
	invalids := validateEnvVars("PT_BATCHER", provided, defined)
	require.ElementsMatch(t, invalids, []string{"PT_BATCHER_FAKE=false"})
}
