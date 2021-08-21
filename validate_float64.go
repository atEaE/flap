package valigo

import "strconv"

// float64Validator is float64 validator.
type float64Validator struct {
	name string
	ptr  *float64
	list []func() error
}

var _ Validator = &float64Validator{} // interface assertion.

// Required means that the value must be entered.
func (v *float64Validator) Required() *float64Validator {
	f := func() error {
		if v.ptr == nil {
			return newRequiredError(v.name)
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Max sets the upper limit.
func (v *float64Validator) Max(limit float64) *float64Validator {
	f := func() error {
		if limit < *v.ptr {
			return newMaxValueOverError(v.name, strconv.FormatFloat(limit, 'f', -1, 64))
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Min sets the lower limit.
func (v *float64Validator) Min(limit float64) *float64Validator {
	f := func() error {
		if limit > *v.ptr {
			return newMinValueOverError(v.name, strconv.FormatFloat(limit, 'f', -1, 64))
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Valid evaluates the validity of the target in turn.
func (v *float64Validator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
