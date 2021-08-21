package valigo_test

import (
	"fmt"
	"testing"

	"github.com/atEaE/valigo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringSliceValidatorRequired(t *testing.T) {
	t.Run("required allow empty slice", func(t *testing.T) {
		// setup
		testcases := []struct {
			name  string
			value *[]string
			want  error
		}{
			{name: "case1", value: stringSlicePtr([]string{}), want: nil},
			{name: "case2", value: stringSlicePtr([]string{"test1", "test2", "test3"}), want: nil},
			{name: "case3", value: stringSlicePtr([]string{""}), want: nil},
			{name: "case4", value: stringSlicePtr([]string{" "}), want: nil},
			{name: "case5", value: nil, want: fmt.Errorf("'case5' is required")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.StringSliceVarP(tc.value, tc.name).Required()

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

	t.Run("required denied empty slice", func(t *testing.T) {
		// setup
		testcases := []struct {
			name  string
			value *[]string
			want  error
		}{
			{name: "case1", value: stringSlicePtr([]string{}), want: fmt.Errorf("'case1' is required")},
			{name: "case2", value: stringSlicePtr([]string{"test1", "test2", "test3"}), want: nil},
			{name: "case3", value: stringSlicePtr([]string{""}), want: nil},
			{name: "case4", value: stringSlicePtr([]string{" "}), want: nil},
			{name: "case5", value: nil, want: fmt.Errorf("'case5' is required")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.StringSliceVarP(tc.value, tc.name, valigo.DeniedEmptySlice()).Required()

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
