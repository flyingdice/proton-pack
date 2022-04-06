package collector

import (
	"errors"
	"github.com/flyingdice/proton-pack/external/kafka/producer"
	"testing"
	"testing/quick"
)

// TestValidation_NewCollector checks that default validation checks are run.
func TestValidation_NewCollector(t *testing.T) {
	checker := func() bool {
		p, err := producer.NewProducer()
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
