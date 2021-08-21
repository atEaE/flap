package valigo

// stringSliceValidator :
type stringSliceValidator struct {
	name            string
	ptr             *[]string
	allowEmptySlice bool
	list            []func() error
}

// StringSliceOption is stringSliceValidator option
type StringSliceOption func(*stringSliceValidator)

func DeniedEmptySlice() StringSliceOption {
	return func(v *stringSliceValidator) {
		v.allowEmptySlice = false
	}
}

var _ Validator = &stringValidator{} // interface assertion.

// Required means that the value must be entered.
func (v *stringSliceValidator) Required() *stringSliceValidator {
	f := func() error {
		if v.ptr == nil {
			return newRequiredError(v.name)
		}

		if !v.allowEmptySlice {
			if len(*v.ptr) == 0 {
				return newRequiredError(v.name)
			}
		}
		return nil
	}
	v.list = append(v.list, f)
	return v
}

// Valid evaluates the validity of the target in turn.
func (v *stringSliceValidator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
