package context

import (
	"context"
	"fmt"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/types/kafka/headers"
	"github.com/flyingdice/proton-pack/internal/types/streams/clock"
	"github.com/flyingdice/proton-pack/internal/types/streams/clock/frozen"
	"github.com/flyingdice/proton-pack/internal/types/streams/clock/standard"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/metadata"
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
	Context  context.Context   `json:"-"`
	Metadata metadata.Metadata `json:"metadata"`
	Headers  headers.Headers   `json:"headers"`
	Clock    clock.Clock       `json:"-"`
}

// NewContext creates and validates a new Context from the given metadata/headers.
func NewContext(ctx context.Context, m metadata.Metadata, h headers.Headers) (Context, validation.ErrorGroup) {
	c := Context{ctx, m, h, standard.Clock{}}
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
		Context:  context.TODO(),
		Metadata: metadata.Generate(rand),
		Headers:  headers.Generate(rand),
		Clock:    frozen.Generate(rand),
	}
}
