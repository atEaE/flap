package valigo

import (
	"os"
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
		if _, err := os.Stat(*v.ptr); os.IsNotExist(err) {
			return newDoesNotExistsError(v.name, *v.ptr)
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// ExpectedDir is the expected value that the target file is a directory.
func (v *fileValidator) ExpectedDir() *fileValidator {
	f := func() error {
		stat, err := os.Stat(*v.ptr)
		if os.IsNotExist(err) {
			return newDoesNotExistsError(v.name, *v.ptr)
		}
		if !stat.IsDir() {
			return newNotDirError(v.name, *v.ptr)
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// ExpectedFile is the expected value that the target file is a file.
func (v *fileValidator) ExpectedFile() *fileValidator {
	f := func() error {
		stat, err := os.Stat(*v.ptr)
		if os.IsNotExist(err) {
			return newDoesNotExistsError(v.name, *v.ptr)
		}
		if stat.IsDir() {
			return newNotFileError(v.name, *v.ptr)
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
