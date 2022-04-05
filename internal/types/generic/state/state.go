package state

import (
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

// State represents a supported state for a Machine.
type State interface {
	~string
}

// GenerateStates returns a slice of random states.
func GenerateStates[T State](rand *rand.Rand) []T {
	faker.SetRandomSource(rand)

	var states []T
	for i := 0; i < rand.Intn(10-1)+1; i++ {
		states = append(states, T(faker.Word()))
	}
	return states
}
