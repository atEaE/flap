package valigo

// stringEnumValidator :
type stringEnumValidator struct {
	name     string
	enum     []string
	ptr      *string
	required bool
	list     []func() error
}

var _ Validator = &stringEnumValidator{} // interface assertion.

// contains will be automatically validated.
func (v *stringEnumValidator) contains() error {
	for i := range v.enum {
		if v.enum[i] == *v.ptr {
			return nil
		}
	}
	return newNotAmongError(v.name, *v.ptr)
}

// Valid evaluates the validity of the target in turn.
func (v *stringEnumValidator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
