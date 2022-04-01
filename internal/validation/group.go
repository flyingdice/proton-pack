package validation

import (
	"fmt"
	"strings"
)

var _ error = (*Errors)(nil)
var _ ErrorGroup = (*Errors)(nil)

// ErrorGroup represents a grouping of validation errors.
type ErrorGroup interface {
	error

	Append(err ...*Error)
	Errors() []*Error
	NilWhenEmpty() ErrorGroup
	Unwrap() error
}

// Errors represents a slice of Error that implements the ErrorGroup interface.
type Errors struct {
	errs []*Error
}

// Append adds a new Error to the group.
//
// Interface: ErrorGroup
func (e *Errors) Append(err ...*Error) {
	e.errs = append(e.errs, err...)
}

// Errors returns a slice of errors in the group.
//
// Interface: ErrorGroup
func (e *Errors) Errors() []*Error {
	return e.errs
}

// NilWhenEmpty will return nil if the instance is nil or doesn't contain
// any errors. This is helpful for callers to return the result of this function
// after accumulating errors in a loop.
//
// Interface: ErrorGroup
func (e *Errors) NilWhenEmpty() ErrorGroup {
	if e == nil {
		return nil
	}
	if len(e.errs) == 0 {
		return nil
	}
	return e
}

// Unwrap returns the next error in the slice or nil if there are no more errors.
//
// Interface: errors.Unwrap, ErrorGroup
func (e *Errors) Unwrap() error {
	// If we have no errors then we do nothing
	if e == nil || len(e.errs) == 0 {
		return nil
	}

	// If we have exactly one error, we can just return that directly.
	if len(e.errs) == 1 {
		return e.errs[0]
	}

	// Shallow copy the slice
	errs := make([]*Error, len(e.errs))
	copy(errs, e.errs)
	return chain(errs)
}

// Error string value of the Errors struct.
//
// Interface: error
func (e *Errors) Error() string {
	if len(e.errs) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, e := range e.errs {
		_, _ = sb.WriteString(fmt.Sprintf("%v\n", e))
	}

	return sb.String()
}
