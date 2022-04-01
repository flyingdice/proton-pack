package headers

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/types/kafka/header"
	"strings"
	"testing"
	"testing/quick"
)

// TestHeaders_String checks String() output is expected format.
func TestHeaders_String(t *testing.T) {
	checker := func(h Headers) bool {
		var b strings.Builder

		for i := 0; i < len(h)-1; i++ {
			_, _ = fmt.Fprintf(&b, "%s, ", h[i])
		}
		if len(h) > 0 {
			_, _ = fmt.Fprintf(&b, "%s", h[len(h)-1])
		}

		return h.String() == fmt.Sprintf("Headers(%s)", b.String())
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestHeaders_EqualsTrue checks equality between two Headers instances.
func TestHeaders_EqualsTrue(t *testing.T) {
	checker := func(h Headers) bool {
		return h.Equals(h)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestHeaders_EqualsFalse checks inequality between two in-equal Headers instances.
func TestHeaders_EqualsFalse(t *testing.T) {
	checker := func(h1 Headers) bool {
		h2 := Headers{}
		h2 = append(h2, header.Header{Key: "foo", Val: []byte("bar")})
		return !h1.Equals(h2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
