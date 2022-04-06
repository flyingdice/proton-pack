package collector

import "github.com/flyingdice/proton-pack/internal/validation"

// ErrProducerMustBeSet is the validation check error returned when
// the collector producer isn't set.
var ErrProducerMustBeSet = validation.NewError(
	"collector_producer_must_be_set",
	"the collector producer must be set and cannot be nil",
)

var defaultChecks = []validation.Check[*Collector]{
	checkProducerNotNil(),
	checkFieldsAreValid(),
}

// checkProducerIsSet validates collector producer is not nil.
func checkProducerNotNil() validation.Check[*Collector] {
	return func(c *Collector) *validation.Error {
		if c.producer == nil {
			return ErrProducerMustBeSet
		}
		return nil
	}
}

// checkFieldsAreValid validates collector fields are valid.
func checkFieldsAreValid() validation.Check[*Collector] {
	return func(c *Collector) *validation.Error {
		if err := validation.RunCheckers(c.machine, c.offsets); err != nil {
			return validation.NewError(
				"collector_must_have_valid_fields",
				err.Error(),
			)
		}
		return nil
	}
}
