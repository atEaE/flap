package valigo

import "strings"

// stringValidator :
type stringValidator struct {
	name            string
	ptr             *string
	allowEmpty      bool
	allowBlankEmpty bool
	list            []func() error
}

type StringOption func(*stringValidator)

func DeniedEmpty() StringOption {
	return func(v *stringValidator) {
		v.allowEmpty = false
	}
}

func DeniedBlankEmpty() StringOption {
	return func(v *stringValidator) {
		v.allowBlankEmpty = false
	}
}

var _ Validator = &stringValidator{} // interface assertion.

// Required means that the value must be entered.
func (v *stringValidator) Required() *stringValidator {
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

// Valid evaluates the validity of the target in turn.
func (v *stringValidator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
