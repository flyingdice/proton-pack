package validation

import "errors"

// chain implements the interfaces necessary for errors.Is/As/Unwrap to
// work in a deterministic way. Is/As/Error will work on the error stored
// in the slice at index zero. Upon an Unwrap call, we will return a chain
// with a new slice with an index shifted by one.
//
// Cribbed from https://github.com/hashicorp/go-multierror
type chain []*Error

// Error implements the error interface
func (e chain) Error() string {
	return e[0].Error()
}

// Unwrap implements errors.Unwrap by returning the next error in the
// chain or nil if there are no more errors.
func (e chain) Unwrap() error {
	if len(e) == 1 {
		return nil
	}
	return e[1:]
}

// As implements errors.As by attempting to map to the current value.
func (e chain) As(target interface{}) bool {
	return errors.As(e[0], target)
}

// Is implements errors.Is by comparing the current value directly.
func (e chain) Is(target error) bool {
	return errors.Is(e[0], target)
}
