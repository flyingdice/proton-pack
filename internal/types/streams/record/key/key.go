package key

import (
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

// Key is the marker interface for generic type constraints for record keys.
// TODO (ahawker) - This will need to include some form of serde support.
type Key interface {
	any
}

// Generate a random Key value.
func Generate(rand *rand.Rand) Key {
	faker.SetRandomSource(rand)
	return faker.Word()
}
