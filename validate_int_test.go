package valigo_test

import (
	"fmt"
	"testing"

	"github.com/atEaE/valigo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntValidatorRequired(t *testing.T) {
	// setup
	testcases := []struct {
		name  string
		value *int
		want  error
	}{
		{name: "case1", value: intPtr(1), want: nil},
		{name: "case2", value: intPtr(213), want: nil},
		{name: "case3", value: nil, want: fmt.Errorf("'case3' is required")},
	}

	for _, tc := range testcases {
		// act
		v := valigo.New()
		v.IntVarP(tc.value, tc.name).Required()

		// assert
		err := v.Validate()
		if tc.want == nil {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			assert.Equal(t, tc.want.Error(), err.Error())
		}
	}
}

func TestIntValidatorMax(t *testing.T) {
	// setup
	testcases := []struct {
		name  string
		value *int
		want  error
	}{
		{name: "case1", value: intPtr(9), want: nil},
		{name: "case2", value: intPtr(10), want: nil},
		{name: "case3", value: intPtr(11), want: fmt.Errorf("value of 'case3' must be less than or equal to 10")},
	}

	for _, tc := range testcases {
		// act
		v := valigo.New()
		v.IntVarP(tc.value, tc.name).Max(10)

		// assert
		err := v.Validate()
		if tc.want == nil {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			assert.Equal(t, tc.want.Error(), err.Error())
		}
	}
}

func TestIntValidatorMin(t *testing.T) {
	// setup
	testcases := []struct {
		name  string
		value *int
		want  error
	}{
		{name: "case1", value: intPtr(6), want: nil},
		{name: "case2", value: intPtr(5), want: nil},
		{name: "case3", value: intPtr(4), want: fmt.Errorf("value of 'case3' must be greater than or equal to 5")},
	}

	for _, tc := range testcases {
		// act
		v := valigo.New()
		v.IntVarP(tc.value, tc.name).Min(5)

		// assert
		err := v.Validate()
		if tc.want == nil {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			assert.Equal(t, tc.want.Error(), err.Error())
		}
	}
}
