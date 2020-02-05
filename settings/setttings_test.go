package settings

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettings(t *testing.T) {
	s := EnvVar("ENV_VAR", "not_set")
	assert.Equal(t, s, "not_set")

	os.Setenv("ENV_VAR", "set")
	s = EnvVar("ENV_VAR", "not_set")
	assert.Equal(t, s, "set")
}
