package valigo

import "strconv"

// int32Validator is integer validator.
type int32Validator struct {
	name string
	ptr  *int32
	list []func() error
}

var _ Validator = &int32Validator{} // interface assertion.

// Required means that the value must be entered.
func (v *int32Validator) Required() *int32Validator {
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
func (v *int32Validator) Max(limit int32) *int32Validator {
	f := func() error {
		if limit < *v.ptr {
			return newMaxValueOverError(v.name, strconv.FormatInt(int64(limit), 10))
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Min sets the lower limit.
func (v *int32Validator) Min(limit int32) *int32Validator {
	f := func() error {
		if limit > *v.ptr {
			return newMinValueOverError(v.name, strconv.FormatInt(int64(limit), 10))
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Valid evaluates the validity of the target in turn.
func (v *int32Validator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
