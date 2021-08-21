package valigo_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/atEaE/valigo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileValidatorRequired(t *testing.T) {
	t.Run("allow empty", func(t *testing.T) {
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
			v.FilepathVarP(tc.value, tc.name).Required()

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

	t.Run("denied empty path", func(t *testing.T) {
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
			v.FilepathVarP(tc.value, tc.name, valigo.DeniedEmptyPath()).Required()

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

	t.Run("denied empty and denied blank empty path", func(t *testing.T) {
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
			v.FilepathVarP(tc.value, tc.name, valigo.DeniedEmptyPath(), valigo.DeniedBlankEmptyPath()).Required()

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

func TestFileValidatorExists(t *testing.T) {

	t.Run("direct filepath", func(t *testing.T) {
		// setup
		testcases := []struct {
			filepath string
			want     error
		}{
			{filepath: "./tests/test.json", want: nil},
			{filepath: "./tests/sample.json", want: fmt.Errorf("'./tests/sample.json' no such file or directory")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.FilepathVar(tc.filepath, "filepath").Exists()

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

	t.Run("wildcard", func(t *testing.T) {
		// setup
		testcases := []struct {
			filepath string
			want     error
		}{
			{filepath: "./tests/file_validator/*.json", want: nil},
			{filepath: "./tests/file_validator/*", want: nil},
			{filepath: "./tests/file_validator/*.go", want: fmt.Errorf("'./tests/file_validator/*.go' no such file or directory")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.FilepathVar(tc.filepath, "filepath").Exists()

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

func TestFileValidatorExistsDir(t *testing.T) {
	t.Run("direct filepath", func(t *testing.T) {
		// setup
		testcases := []struct {
			filepath string
			want     error
		}{
			{filepath: "./tests/file_validator", want: nil},
			{filepath: "./tests/test.json", want: fmt.Errorf("'./tests/test.json' is not directory")},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.FilepathVar(tc.filepath, "filepath").ExistsDir()

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

	t.Run("wildcard", func(t *testing.T) {
		// setup
		testcases := []struct {
			filepath string
			want     error
			winWant  error
		}{
			{filepath: "./tests/file_validator/ok_isdir/*.docker", want: nil, winWant: nil},
			{
				filepath: "./tests/file_validator/ng_isdir/*.docker",
				want:     fmt.Errorf("'tests/file_validator/ng_isdir/2.docker' is not directory"),
				winWant:  fmt.Errorf("'tests\\file_validator\\ng_isdir\\2.docker' is not directory"),
			},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.FilepathVar(tc.filepath, "filepath").ExistsDir()

			// assert
			err := v.Validate()
			if runtime.GOOS == "windows" {
				if tc.winWant == nil {
					require.NoError(t, err)
				} else {
					require.Error(t, err)
					assert.Equal(t, tc.winWant.Error(), err.Error())
				}
			} else {
				if tc.want == nil {
					require.NoError(t, err)
				} else {
					require.Error(t, err)
					assert.Equal(t, tc.want.Error(), err.Error())
				}
			}
		}
	})
}

func TestFileValidatorExistsFile(t *testing.T) {
	t.Run("direct filepath", func(t *testing.T) {
		// setup
		testcases := []struct {
			filepath string
			want     error
		}{
			{filepath: "./tests/file_validator", want: fmt.Errorf("'./tests/file_validator' is not file")},
			{filepath: "./tests/test.json", want: nil},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.FilepathVar(tc.filepath, "filepath").ExistsFile()

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

	t.Run("wildcard", func(t *testing.T) {
		// setup
		testcases := []struct {
			filepath string
			want     error
			winWant  error
		}{
			{filepath: "./tests/file_validator/ok_isfile/*.json", want: nil},
			{
				filepath: "./tests/file_validator/ng_isfile/*.json",
				want:     fmt.Errorf("'tests/file_validator/ng_isfile/test2.json' is not file"),
				winWant:  fmt.Errorf("'tests\\file_validator\\ng_isfile\\test2.json' is not file"),
			},
		}

		for _, tc := range testcases {
			// act
			v := valigo.New()
			v.FilepathVar(tc.filepath, "filepath").ExistsFile()

			// assert
			err := v.Validate()
			if runtime.GOOS == "windows" {
				if tc.winWant == nil {
					require.NoError(t, err)
				} else {
					require.Error(t, err)
					assert.Equal(t, tc.winWant.Error(), err.Error())
				}
			} else {
				if tc.want == nil {
					require.NoError(t, err)
				} else {
					require.Error(t, err)
					assert.Equal(t, tc.want.Error(), err.Error())
				}
			}
		}
	})
}
