package valigo_test

import (
	"fmt"
	"testing"

	"github.com/atEaE/valigo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileValidatorRequired(t *testing.T) {
	t.Run("RequiredAllowEmpty", func(t *testing.T) {
		// setup
		testcases := []struct {
			name  string
			value *string
			want  error
		}{
			{name: "case1", value: stringPtr(""), want: nil},
			{name: "case2", value: stringPtr(" "), want: nil},
			{name: "case3", value: stringPtr("./README.md"), want: nil},
			{name: "case4", value: nil, want: fmt.Errorf("'case4' is required")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.FilepathVarP(tc.value, tc.name).Required(valigo.RequiredAllowEmpty)

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

	t.Run("RequiredDeniedEmpty", func(t *testing.T) {
		// setup
		testcases := []struct {
			name  string
			value *string
			want  error
		}{
			{name: "case1", value: stringPtr(""), want: fmt.Errorf("'case1' is required")},
			{name: "case2", value: stringPtr(" "), want: nil},
			{name: "case3", value: stringPtr("./README.md"), want: nil},
			{name: "case4", value: nil, want: fmt.Errorf("'case4' is required")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.FilepathVarP(tc.value, tc.name).Required(valigo.RequiredDeniedEmpty)

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

	t.Run("DeniedEmptyWithTrimspace", func(t *testing.T) {
		// setup
		testcases := []struct {
			name  string
			value *string
			want  error
		}{
			{name: "case1", value: stringPtr(""), want: fmt.Errorf("'case1' is required")},
			{name: "case2", value: stringPtr(" "), want: fmt.Errorf("'case2' is required")},
			{name: "case3", value: stringPtr("./README.md"), want: nil},
			{name: "case4", value: nil, want: fmt.Errorf("'case4' is required")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.FilepathVarP(tc.value, tc.name).Required(valigo.RequiredDeniedEmptyWithTrimspace)

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
