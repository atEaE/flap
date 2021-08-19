package valigo

// intValidator is integer validator.
type intValidator struct {
	name string
	ptr  *int
	list []func() error
}

var _ Validator = &intValidator{} // interface assertion.

// Required means that the value must be entered.
func (v *intValidator) Required() *intValidator {
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
func (v *intValidator) Max(limit int) *intValidator {
	f := func() error {
		if limit < *v.ptr {
			return newMaxIntValueOverError(v.name, limit)
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Min sets the lower limit.
func (v *intValidator) Min(limit int) *intValidator {
	f := func() error {
		if limit > *v.ptr {
			return newMinIntValueOverError(v.name, limit)
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Valid evaluates the validity of the target in turn.
func (v *intValidator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
