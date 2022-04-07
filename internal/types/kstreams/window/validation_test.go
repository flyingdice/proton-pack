package window

import (
	"github.com/pkg/errors"
	"testing"
	"testing/quick"
	"time"
)

// TestValidation_NewWindow checks that default validation checks are run.
func TestValidation_NewWindow(t *testing.T) {
	checker := func(lo, hi int64) bool {
		_, err := New(time.UnixMicro(lo), time.UnixMicro(hi))
		if lo < 0 {
			if !errors.Is(err, ErrLoMustBeNewerThanUnixEpoch) {
				t.Errorf("expected %v when lo < 0, got %v", ErrLoMustBeNewerThanUnixEpoch, err)
				return false
			}
		} else if hi < lo {
			if !errors.Is(err, ErrHiGreaterThanLo) {
				t.Errorf("expected %v when lo > hi, got %v", ErrHiGreaterThanLo, err)
				return false
			}
		} else {
			if err != nil {
				t.Errorf("expected no errors, got %v", err)
				return false
			}
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
