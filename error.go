package valigo

import "fmt"

// ValidateError
type ValidateError struct {
	name  string
	inner error
}

// Error will output the details.
func (e *ValidateError) Error() string {
	return e.inner.Error()
}

// newValidateError is create error instance.
func newValidateError(name string, err error) *ValidateError {
	return &ValidateError{
		name:  name,
		inner: err,
	}
}

// newRequiredError is create required error instance.
func newRequiredError(name string) *ValidateError {
	return newValidateError(name, fmt.Errorf("'%s' is required", name))
}

// newMaxIntValueOverError is create max intvalue over error instance.
func newMaxIntValueOverError(name string, limit int) *ValidateError {
	return newValidateError(name, fmt.Errorf("value of '%s' must be less than or equal to %d", name, limit))
}

// newMinIntValueOverError is create min intvalue over error instance.
func newMinIntValueOverError(name string, limit int) *ValidateError {
	return newValidateError(name, fmt.Errorf("value of '%s' must be greater than or equal to %d", name, limit))
}

// newDoesNotExistsError is create doesnot exists error instance.
func newDoesNotExistsError(name string, filepath string) *ValidateError {
	return newValidateError(name, fmt.Errorf("'%s' no such file or directory", filepath))
}

// newNotDirError is create not directory error instance.
func newNotDirError(name string, filepath string) *ValidateError {
	return newValidateError(name, fmt.Errorf("'%s' is not directory", filepath))
}

// newNotFileError is create not file error instance.
func newNotFileError(name string, filepath string) *ValidateError {
	return newValidateError(name, fmt.Errorf("'%s' is not file", filepath))
}

// newNotAmongError is create not file error instance.
func newNotAmongError(name string, value interface{}) *ValidateError {
	return newValidateError(name, fmt.Errorf("'%v' is not among the candidates", value))
}
