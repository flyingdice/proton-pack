package context

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/primitive/headers"
	"github.com/flyingdice/proton-pack/internal/record/metadata"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
	"testing/quick"
)

var _ fmt.Stringer = (*Context)(nil)
var _ quick.Generator = (*Context)(nil)
var _ comparison.Equaler = (*Context)(nil)
var _ validation.Checker = (*Context)(nil)

type Context struct {
	Metadata metadata.Metadata `json:"metadata"`
	Headers  headers.Headers   `json:"headers"`
}

// NewContext creates and validates a new Context from the given metadata/headers.
func NewContext(m metadata.Metadata, h headers.Headers) (Context, validation.ErrorGroup) {
	c := Context{m, h}
	return c, c.Check()
}

// Check runs default validation checks for the Context.
func (c Context) Check() validation.ErrorGroup {
	return validation.RunChecks[Context](c, defaultChecks...)
}

// Equals compares two Context instances for equality.
//
// Interface: comparison.Equaler
func (c Context) Equals(v any) bool {
	switch c2 := v.(type) {
	case Context:
		return c.Metadata.Equals(c2.Metadata) && c.Headers.Equals(c2.Headers)
	default:
		return false
	}
}

// Generate random Context values.
//
// Interface: quick.Generator
func (Context) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Context.
//
// Interface: fmt.Stringer.
func (c Context) String() string {
	return fmt.Sprintf(
		"Context(metadata=%s, headers=%s)",
		c.Metadata,
		c.Headers,
	)
}

// Generate a random Context value.
func Generate(rand *rand.Rand) Context {
	return Context{
		Metadata: metadata.Generate(rand),
		Headers:  headers.Generate(rand),
	}
}
