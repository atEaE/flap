package flap

import "fmt"

var (
	// ErrNoSetValidFunc :
	ErrNoSetValidFunc = fmt.Errorf("'ValidFunc' is not set")
)

// flag is base struct
type flag struct {
	Name  string
	Short string
}
