package valigo

import (
	"os"
	"path/filepath"
	"strings"
)

// fileValidator :
type fileValidator struct {
	name string
	ptr  *string
	list []func() error
}

var _ Validator = &fileValidator{} // interface assertion.

// Required means that the value must be entered.
func (v *fileValidator) Required(rt requiredType) *fileValidator {
	f := func() error {
		if v.ptr == nil {
			return newRequiredError(v.name)
		}
		val := *v.ptr
		switch rt {
		case RequiredDeniedEmpty:
			if empty == val {
				return newRequiredError(v.name)
			}
			return nil
		case RequiredDeniedEmptyWithTrimspace:
			if empty == strings.TrimSpace(val) {
				return newRequiredError(v.name)
			}
			return nil
		default: // RequiredAllowEmpty
			return nil
		}
	}
	v.list = append(v.list, f)
	return v
}

// Exists will check for the existence of the target file.
func (v *fileValidator) Exists() *fileValidator {
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
func (v *fileValidator) ExistsDir() *fileValidator {
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
func (v *fileValidator) ExistsFile() *fileValidator {
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
func (v *fileValidator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
