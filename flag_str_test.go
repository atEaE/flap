package flap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestFlagStrValid(t *testing.T) {
	t.Run("set valid test: not set validfunc", func(t *testing.T) {
		// setup
		flag := NewFlagStr("output", "o")

		// act
		err := flag.Valid()

		// assert
		require.Error(t, err)
		assert.Equal(t, ErrNoSetValidFunc, err)
	})

	t.Run("valid call(no error): use option argument", func(t *testing.T) {
		// setup
		validFunc := func(*FlagStr) error { return nil }
		flag := NewFlagStr("output", "o", func(f *FlagStr) { f.ValidFunc = validFunc })

		// act
		err := flag.Valid()

		// assert
		require.NoError(t, err)
	})

	t.Run("valid call(no error): use withValidFunc", func(t *testing.T) {
		// setup
		validFunc := func(*FlagStr) error { return nil }
		flag := NewFlagStr("output", "o").WithValidFunc(validFunc)

		// act
		err := flag.Valid()

		// assert
		require.NoError(t, err)
	})
}
