package flap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlagStrValid(t *testing.T) {
	t.Run("set valid test: not set validfunc", func(t *testing.T) {
		// setup
		flag := FlagStrVar("output", "", "test flags")

		// act
		err := flag.Valid()

		// assert
		require.Error(t, err)
		assert.Equal(t, ErrNoSetValidFunc, err)
	})

	t.Run("valid call(no error): use withValidFunc", func(t *testing.T) {
		// setup
		validFunc := func(*FlagStr) error { return nil }
		flag := FlagStrVar("output", "o", "test flags").WithValidFunc(validFunc)

		// act
		err := flag.Valid()

		// assert
		require.NoError(t, err)
	})
}
