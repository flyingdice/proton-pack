package val

import (
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

// Val is the marker interface for generic type constraints for record values.
// TODO (ahawker) - This will need to include some form of serde support.
type Val interface {
	any
}

// Generate a random Val value.
func Generate(rand *rand.Rand) Val {
	faker.SetRandomSource(rand)
	return faker.Sentence()
}
