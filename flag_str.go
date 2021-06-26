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

// Valid is evaluates the validity of a flag
func (f *FlagStr) Valid() error {
	if f.ValidFunc == nil {
		return ErrNoSetValidFunc
	}
	return f.ValidFunc(f)
}

// WithDefault is to set default value
func (f *FlagStr) WithDefault(def string) *FlagStr {
	f.Default = def
	return f
}

// WithValidFunc is to set validation function
func (f *FlagStr) WithValidFunc(vFunc func(f *FlagStr) error) *FlagStr {
	f.ValidFunc = vFunc
	return f
}
