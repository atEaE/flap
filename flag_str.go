package flap

import (
	"github.com/spf13/pflag"
)

// ValidFuncString is validate function for FlagStr(string)
type ValidFuncString func(*FlagStr) error

// FlagStr is string flag.
type FlagStr struct {
	*pflag.Flag

	value     string
	validFunc ValidFuncString
}

// FlagStrVar is create a new FlagStr instance.
func FlagStrVar(name string, value string, usage string) *FlagStr {
	return FlagStrVarP(name, "", value, usage)
}

/// FlagStrVarP is create a new FlagStr instance
func FlagStrVarP(name, shorthand, value string, usage string) *FlagStr {
	f := &FlagStr{}
	f.Flag = newFlag(newStringValue(value, &f.value), name, shorthand, usage)

	return f
}

// Value is flag value.
func (f *FlagStr) Value() string {
	return f.value
}

// Valid is evaluates the validity of a flag
func (f *FlagStr) Valid() error {
	if f.validFunc == nil {
		return ErrNoSetValidFunc
	}
	return f.validFunc(f)
}

// WithValidFunc is evaluates the validity of a flag
func (f *FlagStr) WithValidFunc(validate ValidFuncString) *FlagStr {
	f.validFunc = validate
	return f
}

type stringValue string

func newStringValue(val string, p *string) *stringValue {
	*p = val
	return (*stringValue)(p)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}
func (s *stringValue) Type() string {
	return "string"
}

func (s *stringValue) String() string { return string(*s) }
