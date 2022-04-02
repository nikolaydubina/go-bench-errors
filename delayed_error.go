package gobencherrors

import "fmt"

type DelayedPrefixError struct {
	Err    error
	Prefix string
	V      float64
}

func (e DelayedPrefixError) Error() string {
	if e.Err == nil {
		return e.Prefix + fmt.Sprintf("%f", e.V)
	}
	return e.Prefix + fmt.Sprintf("%f", e.V) + e.Err.Error()
}
