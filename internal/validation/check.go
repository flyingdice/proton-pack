package validation

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"hash/fnv"
)

// errCode returns a numeric value from the given check error name.
//
// TODO (ahawker) - This simple hash is not guaranteed to be unique.
func errCode(name string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(name))
	return h.Sum32()
}

// CheckError represents the failure of a validation check.
type CheckError struct {
	// Unique numeric value that represents a specific type of check error.
	Code uint32
	// Unique human understandable name of the error.
	Slug string
	// Full description of the error with any additional context.
	Desc string
}

// Return CheckError representation as a string.
//
// Interface: error
func (e *CheckError) Error() string {
	return fmt.Sprintf("CheckError(code=%d slug=%s desc=%s)", e.Code, e.Slug, e.Desc)
}

// NewCheckError creates a new CheckError for the given slug/desc.
func NewCheckError(slug, desc string) *CheckError {
	return &CheckError{
		Code: errCode(slug),
		Slug: slug,
		Desc: desc,
	}
}

// Check is a function that performs validation of a value.
type Check[T any] func(val T) *CheckError

// Validate runs checks for the given value.
//
// If the value is invalid, one or more errors will be returned.
func Validate[T any](val T, checks ...Check[T]) (err error) {
	for _, check := range checks {
		if e := check(val); e != nil {
			err = multierror.Append(err, e)
		}
	}
	return
}
