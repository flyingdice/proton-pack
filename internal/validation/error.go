package validation

import (
	"fmt"
	"hash/fnv"
)

var _ error = (*Error)(nil)

// Error represents the failure of a validation check.
type Error struct {
	// Unique numeric value that represents a specific type of check error.
	Code uint32
	// Unique human understandable name of the error.
	Slug string
	// Full description of the error with any additional context.
	Desc string
}

// Error string value of the Error struct.
//
// Interface: error
func (e *Error) Error() string {
	return fmt.Sprintf("Error(code=%d slug=%s desc=%s)", e.Code, e.Slug, e.Desc)
}

// NewError creates a new Error for the given slug/desc.
func NewError(slug, desc string) *Error {
	return &Error{
		Code: errCode(slug),
		Slug: slug,
		Desc: desc,
	}
}

// errCode returns a numeric value from the given check error name.
//
// TODO (ahawker) - This simple hash is not guaranteed to be unique.
func errCode(name string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(name))
	return h.Sum32()
}
