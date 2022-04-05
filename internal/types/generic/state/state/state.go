package state

import (
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

// State represents a generic type constraint for machine states.
type State interface {
	~string
}

// Generate returns a slice of random states.
func Generate[T State](rand *rand.Rand) []T {
	faker.SetRandomSource(rand)

	var states []T
	for i := 0; i < rand.Intn(10-1)+1; i++ {
		states = append(states, T(faker.Word()))
	}
	return states
}
