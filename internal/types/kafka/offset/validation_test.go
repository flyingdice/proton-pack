package offset

import (
	"github.com/pkg/errors"
	"testing"
	"testing/quick"
)

// TestValidation_NewOffset checks that default validation checks are run.
func TestValidation_NewOffset(t *testing.T) {
	checker := func(raw int64) bool {
		_, err := New(raw)
		if raw < 0 {
			if !errors.Is(err, ErrMustBePositive) {
				t.Errorf("expected %v when offset < 0, got %v", ErrMustBePositive, err)
				return false
			}
		} else {
			if err != nil {
				t.Errorf("expected no error when offset >= 0, got %v", err)
				return false
			}
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
