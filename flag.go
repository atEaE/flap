package flap

import (
	"fmt"

	"github.com/spf13/pflag"
)

var (
	// ErrNoSetValidFunc :
	ErrNoSetValidFunc = fmt.Errorf("'validFunc' is not set")
)

// newFlag is create a new pflag.Flag instance.
func newFlag(value pflag.Value, name, shorthand, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:      name,
		Shorthand: shorthand,
		Usage:     usage,
		Value:     value,
		DefValue:  value.String(),
	}
}
