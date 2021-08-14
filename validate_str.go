package valigo

// stringValidator :
type stringValidator struct {
	name string
	ptr  *string
	list []func() error
}

var _ Validator = &stringValidator{} // interface assertion.

// Required means that the value must be entered.
func (v *stringValidator) Required() *stringValidator {
	f := func() error {
		if v.ptr == nil {
			return newRequiredError(v.name)
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
