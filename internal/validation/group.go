package validation

import (
	"fmt"
	"strings"
)

var _ error = (*errorGroup)(nil)

// ErrorGroup represents a grouping of validation errors.
type ErrorGroup interface {
	error

	Append(err *Error)
	Unwrap() error
	NilWhenEmpty() ErrorGroup
}

// errorGroup represents a slice of Error.
type errorGroup struct {
	Errors []*Error
}

// NilWhenEmpty will return nil if the instance is nil or doesn't contain
// any errors. This is helpful for callers to return the result of this function
// after accumulating errors in a loop.
func (e *errorGroup) NilWhenEmpty() ErrorGroup {
	if e == nil {
		return nil
	}
	if len(e.Errors) == 0 {
		return nil
	}
	return e
}

// Append adds a new Error to the group.
func (e *errorGroup) Append(err *Error) {
	e.Errors = append(e.Errors, err)
}

// Unwrap returns the next error in the slice or nil if there are no more errors.
//
// Interface: errors.Unwrap
func (e *errorGroup) Unwrap() error {
	// If we have no errors then we do nothing
	if e == nil || len(e.Errors) == 0 {
		return nil
	}

	// If we have exactly one error, we can just return that directly.
	if len(e.Errors) == 1 {
		return e.Errors[0]
	}

	// Shallow copy the slice
	errs := make([]*Error, len(e.Errors))
	copy(errs, e.Errors)
	return chain(errs)
}

// Error string value of the errorGroup struct.
//
// Interface: error
func (e *errorGroup) Error() string {
	if len(e.Errors) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, e := range e.Errors {
		_, err := sb.WriteString(fmt.Sprintf("%v\n", e))
		if err != nil {
			panic(err)
		}
	}

	return sb.String()
}
