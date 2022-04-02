package context

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestContext_String checks String() output is expected format.
func TestContext_String(t *testing.T) {
	checker := func(c Context) bool {
		return c.String() == fmt.Sprintf(
			"Context(metadata=%s, headers=%v)",
			c.Metadata,
			c.Headers,
		)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
