package store

import (
	"fmt"
	"testing"
	"testing/quick"
)

// TestStore_String checks String() output is expected format.
func TestStore_String(t *testing.T) {
	checker := func(s *Store) bool {
		return s.String() == fmt.Sprintf("Store(len=%d)", len(s.store))
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
