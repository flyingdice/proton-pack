package header

import (
	"github.com/pkg/errors"
	"testing"
	"testing/quick"
)

// TestValidation_NewHeader checks that default validation checks are run.
func TestValidation_NewHeader(t *testing.T) {
	checker := func(key string, val []byte) bool {
		_, err := New(key, val)
		if key == "" {
			if !errors.Is(err, ErrKeyMustBeSet) {
				t.Errorf("expected %v when header is empty string, got %v", ErrKeyMustBeSet, err)
				return false
			}
		} else {
			if err != nil {
				t.Errorf("expected no error when header >= 0, got %v", err)
				return false
			}
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
