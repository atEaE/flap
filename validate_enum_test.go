package valigo_test

import (
	"fmt"
	"testing"

	"github.com/atEaE/valigo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestType string

var (
	TestTypeSample1 = TestType("sample1")
	TestTypeSample2 = TestType("sample2")
)

func TestEnumValidatorContainsPrimitive(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		// setup
		testcases := []struct {
			title string
			value interface{}
			enums []interface{}
			want  error
		}{
			{title: "ok", value: "test", enums: []interface{}{"sample", "example", "test"}, want: nil},
			{title: "ng", value: "test", enums: []interface{}{"sample", "example"}, want: fmt.Errorf("'test' is not among the candidates")},
		}

		for _, tc := range testcases {
			t.Run(tc.title, func(t *testing.T) {
				// act
				v := valigo.New()
				v.EnumVar(tc.value, "value", tc.enums)

				// assert
				err := v.Validate()
				if tc.want == nil {
					require.NoError(t, err)
				} else {
					require.Error(t, err)
					assert.Equal(t, tc.want.Error(), err.Error())
				}
			})
		}
	})

	t.Run("int", func(t *testing.T) {
		// setup
		testcases := []struct {
			title string
			value interface{}
			enums []interface{}
			want  error
		}{
			{title: "ok", value: 2, enums: []interface{}{1, 2, 3}, want: nil},
			{title: "ng", value: 2, enums: []interface{}{1, 3, 4}, want: fmt.Errorf("'2' is not among the candidates")},
		}

		for _, tc := range testcases {
			t.Run(tc.title, func(t *testing.T) {
				// act
				v := valigo.New()
				v.EnumVar(tc.value, "value", tc.enums)

				// assert
				err := v.Validate()
				if tc.want == nil {
					require.NoError(t, err)
				} else {
					require.Error(t, err)
					assert.Equal(t, tc.want.Error(), err.Error())
				}
			})
		}
	})

	t.Run("float64", func(t *testing.T) {
		// setup
		var testValue float64 = 2.1
		okEnums := []float64{1.6, 2.1, 3.3}
		ngEnums := []float64{1.6, 3.3, 4.8}

		testcases := []struct {
			title string
			value interface{}
			enums []interface{}
			want  error
		}{
			{title: "ok", value: testValue, enums: []interface{}{okEnums[0], okEnums[1], okEnums[2]}, want: nil},
			{title: "ng", value: testValue, enums: []interface{}{ngEnums[0], ngEnums[1], ngEnums[2]}, want: fmt.Errorf("'2.1' is not among the candidates")},
		}

		for _, tc := range testcases {
			t.Run(tc.title, func(t *testing.T) {
				// act
				v := valigo.New()
				v.EnumVar(tc.value, "value", tc.enums)

				// assert
				err := v.Validate()
				if tc.want == nil {
					require.NoError(t, err)
				} else {
					require.Error(t, err)
					assert.Equal(t, tc.want.Error(), err.Error())
				}
			})
		}
	})
}
