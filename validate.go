package valigo

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
	valid := &stringValidator{name: name, ptr: stringPtr(arg)}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) StringVarP(arg *string, name string) *stringValidator {
	valid := &stringValidator{name: name, ptr: arg}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) FilapathVar(arg string, name string) *fileValidator {
	valid := &fileValidator{name: name, ptr: stringPtr(arg)}
	v.list = append(v.list, valid)
	return valid
}

func (v *Valigo) FilapathVarP(arg *string, name string) *fileValidator {
	valid := &fileValidator{name: name, ptr: arg}
	v.list = append(v.list, valid)
	return valid
}
