package valigo_test

import (
	"fmt"
	"testing"

	"github.com/atEaE/valigo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringValidatorRequired(t *testing.T) {
	t.Run("allow empty", func(t *testing.T) {
		// setup
		testcases := []struct {
			name  string
			value *string
			want  error
		}{
			{name: "case1", value: stringPtr(""), want: nil},
			{name: "case2", value: stringPtr(" "), want: nil},
			{name: "case3", value: stringPtr("test_string"), want: nil},
			{name: "case4", value: nil, want: fmt.Errorf("'case4' is required")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.StringVarP(tc.value, tc.name).Required()

			// assert
			err := v.Validate()
			if tc.want == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Equal(t, tc.want.Error(), err.Error())
			}
		}
	})

	t.Run("denied empty", func(t *testing.T) {
		// setup
		testcases := []struct {
			name  string
			value *string
			want  error
		}{
			{name: "case1", value: stringPtr(""), want: fmt.Errorf("'case1' is required")},
			{name: "case2", value: stringPtr(" "), want: nil},
			{name: "case3", value: stringPtr("test_string"), want: nil},
			{name: "case4", value: nil, want: fmt.Errorf("'case4' is required")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.StringVarP(tc.value, tc.name, valigo.DeniedEmpty()).Required()

			// assert
			err := v.Validate()
			if tc.want == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Equal(t, tc.want.Error(), err.Error())
			}
		}
	})

	t.Run("denied empty and denied blank empty", func(t *testing.T) {
		// setup
		testcases := []struct {
			name  string
			value *string
			want  error
		}{
			{name: "case1", value: stringPtr(""), want: fmt.Errorf("'case1' is required")},
			{name: "case2", value: stringPtr(" "), want: fmt.Errorf("'case2' is required")},
			{name: "case3", value: stringPtr("test_string"), want: nil},
			{name: "case4", value: nil, want: fmt.Errorf("'case4' is required")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.StringVarP(tc.value, tc.name, valigo.DeniedEmpty(), valigo.DeniedBlankEmpty()).Required()

			// assert
			err := v.Validate()
			if tc.want == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Equal(t, tc.want.Error(), err.Error())
			}
		}
	})
}
