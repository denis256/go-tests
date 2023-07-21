package config_test

import (
	"noerrorcheck/config"
	"testing"

)

func TestConfigCheck(t *testing.T) {
	t.Parallel()

	ctx, _ := config.CreateEvalConfig("qwe")
	assert.Equal(t, "qwe", ctx)
}
