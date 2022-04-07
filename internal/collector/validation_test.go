package collector

import (
	"github.com/flyingdice/proton-pack/external/kafka/producer"
	"github.com/pkg/errors"
	"testing"
	"testing/quick"
)

// TestValidation_NewCollector checks that default validation checks are run.
func TestValidation_NewCollector(t *testing.T) {
	checker := func() bool {
		p, err := producer.New()
		if err != nil {
			t.Error(err.Error())
			return false
		}
		_, err = New(p)
		if p == nil {
			if !errors.Is(err, ErrProducerMustBeSet) {
				t.Errorf("expected %v when producer is nil, got %v", ErrProducerMustBeSet, err)
				return false
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error %v", err)
				return false
			}
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
