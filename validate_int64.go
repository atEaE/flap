package valigo

import "strconv"

// int64Validator is integer validator.
type int64Validator struct {
	name string
	ptr  *int64
	list []func() error
}

var _ Validator = &int64Validator{} // interface assertion.

// Required means that the value must be entered.
func (v *int64Validator) Required() *int64Validator {
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
func (v *int64Validator) Max(limit int64) *int64Validator {
	f := func() error {
		if limit < *v.ptr {
			return newMaxValueOverError(v.name, strconv.FormatInt(limit, 10))
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Min sets the lower limit.
func (v *int64Validator) Min(limit int64) *int64Validator {
	f := func() error {
		if limit > *v.ptr {
			return newMinValueOverError(v.name, strconv.FormatInt(limit, 10))
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Valid evaluates the validity of the target in turn.
func (v *int64Validator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
