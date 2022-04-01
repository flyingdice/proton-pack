package headers

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/flyingdice/proton-pack/internal/comparison"
	"github.com/flyingdice/proton-pack/internal/primitive/header"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
	"strings"
	"testing/quick"
)

var _ fmt.Stringer = (*Headers)(nil)
var _ quick.Generator = (*Headers)(nil)
var _ comparison.Equaler = (*Headers)(nil)
var _ validation.Checker = (*Headers)(nil)

// Headers represents an ordered slice of header.Header structs.
type Headers []header.Header

// NewHeaders creates and validates new Headers from the given slice of header.Header.
func NewHeaders(headers []header.Header) (Headers, validation.ErrorGroup) {
	h := Headers(headers)
	return h, h.Check()
}

// Check runs default validation checks for the Header.
func (h Headers) Check() validation.ErrorGroup {
	return validation.Validate[Headers](h, defaultChecks...)
}

// Equals compares two Headers instances for equality.
//
// Interface: comparison.Equaler
func (h Headers) Equals(v any) bool {
	switch h2 := v.(type) {
	case Headers:
		if len(h) != len(h2) {
			return false
		}
		for i := range h {
			if !h[i].Equals(h2[i]) {
				return false
			}
		}
		return true
	case []header.Header:
		return h.Equals(Headers(h2))
	default:
		return false
	}
}

// Generate random Headers values.
//
// Interface: quick.Generator
func (Headers) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand))
}

// String value of the Headers.
//
// Interface: fmt.Stringer.
func (h Headers) String() string {
	var b strings.Builder

	for i := 0; i < len(h)-1; i++ {
		_, _ = fmt.Fprintf(&b, "%s, ", h[i])
	}
	if len(h) > 0 {
		_, _ = fmt.Fprintf(&b, "%s", h[len(h)-1])
	}

	return fmt.Sprintf("Headers(%s)", b.String())
}

// Generate a random Headers value.
func Generate(rand *rand.Rand) Headers {
	faker.SetRandomSource(rand)

	var headers []header.Header
	for i := 0; i < rand.Intn(10); i++ {
		headers = append(headers, header.Generate(rand))
	}
	return headers
}
