package config

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("should_return_an_errParseEnv_error", func(t *testing.T) {
		_, err := New()

		assert.Equal(t, true, errors.Is(err, errParseEnv))
	})

	t.Run("should_return_an_errParseEnv_error_because_validator_failed", func(t *testing.T) {
		t.Setenv("PORT", "test")
		t.Setenv("ENV", "test")
		t.Setenv("REDIS_HOST", "test")
		t.Setenv("REDIS_PORT", "test")

		_, err := New()
		assert.Equal(t, true, errors.Is(err, errParseEnv))
	})

	t.Run("shoul_ be_ok", func(t *testing.T) {
		t.Setenv("PORT", "v")
		t.Setenv("ENV", "debug")
		t.Setenv("REDIS_HOST", "v")
		t.Setenv("REDIS_PORT", "v")

		c, err := New()

		assert.NoError(t, err)
		assert.Equal(t, Conf{"v", "debug", "v", "v"}, c)
	})
}
