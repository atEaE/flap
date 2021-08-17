package valigo

import (
	"reflect"
)

const (
	empty = ""
)

// Validator is
type Validator interface {
	Valid() error
}

// Valigo is validation managements struct.
type Valigo struct {
	list []Validator
}

// New returns a new valigo instance.
func New() *Valigo {
	return &Valigo{}
}

// Validate evaluates the validity of the target in turn.
func (v *Valigo) Validate() error {
	for _, f := range v.list {
		if err := f.Valid(); err != nil {
			return err
		}
	}
	return nil
}

func (v *Valigo) StringVar(arg string, name string) *stringValidator {
	return v.StringVarP(&arg, name)
}

func (v *Valigo) StringVarP(arg *string, name string) *stringValidator {
	valid := &stringValidator{name: name, ptr: arg}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) EnumVar(arg interface{}, name string, enums []interface{}) *enumValidator {
	val := reflect.ValueOf(arg)
	if val.Kind() == reflect.Ptr {
		panic("if you want to use a pointer as a variable, use 'EnumVarP'")
	}

	for i := range enums {
		if val.Type() != reflect.TypeOf(enums[i]) {
			panic("type of the enumeration candidate is different from the type of the variable")
		}
	}
	valid := &enumValidator{
		name:  name,
		value: arg,
		vType: val.Type(),
		enums: enums,
		isPtr: false,
	}
	valid.list = append(valid.list, valid.contains)
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) EnumVarP(arg interface{}, name string, enums []interface{}) *enumValidator {
	valP := reflect.ValueOf(arg)
	if valP.Kind() != reflect.Ptr {
		panic("if you want to use a variable, use 'EnumVar'")
	}
	refValType := valP.Type().Elem()

	for i := range enums {
		if refValType != reflect.TypeOf(enums[i]) {
			panic("type of the enumeration candidate is different from the type of the variable")
		}
	}
	valid := &enumValidator{
		name:  name,
		value: arg,
		vType: refValType,
		enums: enums,
		isPtr: true,
	}
	valid.list = append(valid.list, valid.contains)
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) FilepathVar(arg string, name string) *fileValidator {
	return v.FilepathVarP(&arg, name)
}

func (v *Valigo) FilepathVarP(arg *string, name string) *fileValidator {
	valid := &fileValidator{name: name, ptr: arg}
	v.list = append(v.list, valid)
	return valid
}
