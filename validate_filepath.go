package valigo

import (
	"os"
	"path/filepath"
	"strings"
)

// filepathValidator :
type filepathValidator struct {
	name            string
	ptr             *string
	allowEmpty      bool
	allowBlankEmpty bool
	list            []func() error
}

type FilepathOption func(*filepathValidator)

func DeniedEmptyPath() FilepathOption {
	return func(v *filepathValidator) {
		v.allowEmpty = false
	}
}

func DeniedBlankEmptyPath() FilepathOption {
	return func(v *filepathValidator) {
		v.allowBlankEmpty = false
	}
}

var _ Validator = &filepathValidator{} // interface assertion.

// Required means that the value must be entered.
func (v *filepathValidator) Required() *filepathValidator {
	f := func() error {
		if v.ptr == nil {
			return newRequiredError(v.name)
		}
		if !v.allowEmpty {
			if empty == *v.ptr {
				return newRequiredError(v.name)
			}
		}
		if !v.allowBlankEmpty {
			if empty == strings.TrimSpace(*v.ptr) {
				return newRequiredError(v.name)
			}
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Exists will check for the existence of the target file.
func (v *filepathValidator) Exists() *filepathValidator {
	f := func() error {
		files, err := filepath.Glob(*v.ptr)
		if err != nil || len(files) == 0 {
			return newDoesNotExistsError(v.name, *v.ptr)
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// ExistsDir is the expected value that the target file is a directory.
func (v *filepathValidator) ExistsDir() *filepathValidator {
	f := func() error {
		files, err := filepath.Glob(*v.ptr)
		if err != nil || len(files) == 0 {
			return newDoesNotExistsError(v.name, *v.ptr)
		}
		for _, f := range files {
			stat, err := os.Stat(f)
			if os.IsNotExist(err) {
				return newDoesNotExistsError(v.name, f)
			}
			if !stat.IsDir() {
				return newNotDirError(v.name, f)
			}
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// ExistsFile is the expected value that the target file is a file.
func (v *filepathValidator) ExistsFile() *filepathValidator {
	f := func() error {
		files, err := filepath.Glob(*v.ptr)
		if err != nil || len(files) == 0 {
			return newDoesNotExistsError(v.name, *v.ptr)
		}
		for _, f := range files {
			stat, err := os.Stat(f)
			if os.IsNotExist(err) {
				return newDoesNotExistsError(v.name, f)
			}
			if stat.IsDir() {
				return newNotFileError(v.name, f)
			}
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Valid evaluates the validity of the target in turn.
func (v *filepathValidator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
