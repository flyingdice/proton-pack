package topic

import (
	"errors"
	"regexp"
	"testing"
	"testing/quick"
)

// TestValidation_NewTopic checks that default validation checks are run.
func TestValidation_NewTopic(t *testing.T) {
	regex := regexp.MustCompile(`[a-zA-Z0-9_.\-]{1,255}`)

	checker := func(raw string) bool {
		_, err := New(raw)
		if !regex.MatchString(raw) {
			if !errors.Is(err, ErrMustMatchPattern) {
				t.Errorf("expected %v when topic doesn't match pattern, got %v", ErrMustMatchPattern, err)
				return false
			}
		} else {
			if err != nil {
				t.Errorf("expected no error when topic matches pattern, got %v", err)
				return false
			}
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
