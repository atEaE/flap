package flap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagStrSetDefault(t *testing.T) {
	t.Run("set default test: not use default", func(t *testing.T) {
		// setup
		flag := NewFlagStr("output", "o")

		// act & assert
		assert.Equal(t, "", flag.Default)
	})

	t.Run("set default test: use option arguments", func(t *testing.T) {
		// setup
		flag := NewFlagStr("output", "o", func(f *FlagStr) { f.Default = "output.txt" })

		// act & assert
		assert.Equal(t, "output.txt", flag.Default)
	})

	t.Run("set default test: use withDefault", func(t *testing.T) {
		// setup
		flag := NewFlagStr("output", "o").WithDefault("output.txt")

		// act & assert
		assert.Equal(t, "output.txt", flag.Default)
	})
}
