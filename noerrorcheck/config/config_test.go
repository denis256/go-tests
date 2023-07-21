package config_test

import (
	"testing"

	assert "github.com/stretchr/testify/assert"

	"noerrorcheck/config"
)

func TestConfigCheck(t *testing.T) {
	t.Parallel()

	ctx, _ := config.CreateEvalConfig("qwe")
	assert.Equal(t, "qwe", ctx)
}
