package valigo

import "reflect"

// stringEnumValidator :
type enumValidator struct {
	name  string
	enums []interface{}
	value interface{}
	isPtr bool
	vType reflect.Type
	list  []func() error
}

var _ Validator = &enumValidator{} // interface assertion.

// contains will be automatically validated.
func (v *enumValidator) contains() error {
	if v.isPtr {
		refVal := reflect.ValueOf(v.value).Elem().Interface()
		for i := range v.enums {
			if v.enums[i] == refVal {
				return nil
			}
		}
		return newNotAmongError(v.name, refVal)
	} else {
		for i := range v.enums {
			if v.enums[i] == v.value {
				return nil
			}
		}
		return newNotAmongError(v.name, v.value)
	}
}

// Valid evaluates the validity of the target in turn.
func (v *enumValidator) Valid() error {
	for _, f := range v.list {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}
