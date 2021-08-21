package valigo

import "strconv"

// float32Validator is float32 validator.
type float32Validator struct {
	name string
	ptr  *float32
	list []func() error
}

var _ Validator = &float32Validator{} // interface assertion.

// Required means that the value must be entered.
func (v *float32Validator) Required() *float32Validator {
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
func (v *float32Validator) Max(limit float32) *float32Validator {
	f := func() error {
		if limit < *v.ptr {
			return newMaxValueOverError(v.name, strconv.FormatFloat(float64(limit), 'f', -1, 32))
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Min sets the lower limit.
func (v *float32Validator) Min(limit float32) *float32Validator {
	f := func() error {
		if limit > *v.ptr {
			return newMinValueOverError(v.name, strconv.FormatFloat(float64(limit), 'f', -1, 32))
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Valid evaluates the validity of the target in turn.
func (v *float32Validator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
