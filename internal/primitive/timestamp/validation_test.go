package timestamp

import (
	"errors"
	"testing"
	"testing/quick"
	"time"
)

func TestValidation_NewOffset(t *testing.T) {
	checker := func(raw int64) bool {
		_, err := NewTimestamp(time.UnixMilli(raw))
		if raw < 0 {
			if !errors.Is(err, ErrMustBeNewerThanUnixEpoch) {
				t.Errorf("expected %v when timestamp < 0, got %v", ErrMustBeNewerThanUnixEpoch, err)
				return false
			}
		} else {
			if err != nil {
				t.Errorf("expected no error when timestamp >= 0, got %v", err)
				return false
			}
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
