package valigo_test

import (
	"fmt"
	"testing"

	"github.com/atEaE/valigo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFloat32ValidatorRequired(t *testing.T) {
	// setup
	testcases := []struct {
		name  string
		value *float32
		want  error
	}{
		{name: "case1", value: float32Ptr(1.1), want: nil},
		{name: "case2", value: float32Ptr(213.457), want: nil},
		{name: "case3", value: nil, want: fmt.Errorf("'case3' is required")},
	}

	for _, tc := range testcases {
		// act
		v := valigo.New()
		v.Float32VarP(tc.value, tc.name).Required()

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

func TestFloat32ValidatorMax(t *testing.T) {
	// setup
	testcases := []struct {
		name  string
		value float32
		want  error
	}{
		{name: "case1", value: 10.19, want: nil},
		{name: "case2", value: 10.20, want: nil},
		{name: "case3", value: 10.21, want: fmt.Errorf("value of 'case3' must be less than or equal to 10.2")},
	}

	for _, tc := range testcases {
		// act
		v := valigo.New()
		v.Float32Var(tc.value, tc.name).Max(10.20)

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

func TestFloat32ValidatorMin(t *testing.T) {
	// setup
	testcases := []struct {
		name  string
		value float32
		want  error
	}{
		{name: "case1", value: 10.21, want: nil},
		{name: "case2", value: 10.20, want: nil},
		{name: "case3", value: 10.19, want: fmt.Errorf("value of 'case3' must be greater than or equal to 10.2")},
	}

	for _, tc := range testcases {
		// act
		v := valigo.New()
		v.Float32Var(tc.value, tc.name).Min(10.20)

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
