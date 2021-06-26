package flap

// FlagStr is string flag.
type FlagStr struct {
	flag

	Default   string
	Value     string
	ValidFunc func(f *FlagStr) error
}

// NewFlagStr creates a new FlagStr instance.
func NewFlagStr(name, short string, opts ...func(*FlagStr)) *FlagStr {
	flag := &FlagStr{
		flag: flag{
			Name:  name,
			Short: short,
		},
	}
	for _, opt := range opts {
		opt(flag)
	}
	return flag
}

// WithDefault is to set default value
func (f *FlagStr) WithDefault(def string) *FlagStr {
	f.Default = def
	return f
}
