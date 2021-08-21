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

func (v *Valigo) StringVar(arg string, name string, opt ...StringOption) *stringValidator {
	return v.StringVarP(&arg, name, opt...)
}

func (v *Valigo) StringVarP(arg *string, name string, opt ...StringOption) *stringValidator {
	valid := &stringValidator{name: name, ptr: arg, allowEmpty: true, allowBlankEmpty: true}
	for _, o := range opt {
		o(valid)
	}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) StringSliceVar(arg []string, name string, opt ...StringSliceOption) *stringSliceValidator {
	return v.StringSliceVarP(&arg, name, opt...)
}

func (v *Valigo) StringSliceVarP(arg *[]string, name string, opt ...StringSliceOption) *stringSliceValidator {
	valid := &stringSliceValidator{name: name, ptr: arg, allowEmptySlice: true}
	for _, o := range opt {
		o(valid)
	}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) IntVar(arg int, name string) *intValidator {
	return v.IntVarP(&arg, name)
}

func (v *Valigo) IntVarP(arg *int, name string) *intValidator {
	valid := &intValidator{name: name, ptr: arg}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) Int64Var(arg int64, name string) *int64Validator {
	return v.Int64VarP(&arg, name)
}

func (v *Valigo) Int64VarP(arg *int64, name string) *int64Validator {
	valid := &int64Validator{name: name, ptr: arg}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) Int32Var(arg int32, name string) *int32Validator {
	return v.Int32VarP(&arg, name)
}

func (v *Valigo) Int32VarP(arg *int32, name string) *int32Validator {
	valid := &int32Validator{name: name, ptr: arg}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) Float64Var(arg float64, name string) *float64Validator {
	return v.Float64VarP(&arg, name)
}

func (v *Valigo) Float64VarP(arg *float64, name string) *float64Validator {
	valid := &float64Validator{name: name, ptr: arg}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) Float32Var(arg float32, name string) *float32Validator {
	return v.Float32VarP(&arg, name)
}

func (v *Valigo) Float32VarP(arg *float32, name string) *float32Validator {
	valid := &float32Validator{name: name, ptr: arg}
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

func (v *Valigo) FilepathVar(arg string, name string, opt ...FilepathOption) *filepathValidator {
	return v.FilepathVarP(&arg, name, opt...)
}

func (v *Valigo) FilepathVarP(arg *string, name string, opt ...FilepathOption) *filepathValidator {
	valid := &filepathValidator{name: name, ptr: arg, allowEmpty: true, allowBlankEmpty: true}
	for _, o := range opt {
		o(valid)
	}
	v.list = append(v.list, valid)
	return valid
}
