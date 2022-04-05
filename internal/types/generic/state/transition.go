package state

import (
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"reflect"
)

var Valid = struct{}{}
var Invalid = struct{}{}

// Transitions represents a structure for describing states and
// valid transitions between them.
type Transitions[T State] struct {
	transitions map[T]map[T]struct{}
}

// Valid return true for registered transition between current -> next states.
func (t *Transitions[T]) Valid(current, next T) bool {
	c, ok := t.transitions[current]
	if !ok {
		return false
	}
	n, ok := c[next]
	if !ok {
		return false
	}
	return n == Valid
}

// Len returns number of registered transitions for all states.
func (t *Transitions[T]) Len() int {
	return len(t.transitions)
}

// Generate random Transition values.
//
// Interface: quick.Generator
func (Transitions[T]) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(GenerateTransitions[T](rand, []T{}))
}

// GenerateTransitions from a set of random states.
func GenerateTransitions[T State](rand *rand.Rand, states []T) Transitions[T] {
	faker.SetRandomSource(rand)

	if len(states) == 0 {
		states = GenerateStates[T](rand)
	}

	// Register random transitions for states within the given set.
	transitions := make(map[T]map[T]struct{})
	for i := 0; i < rand.Intn(len(states)); i++ {
		registered := make(map[T]struct{})
		for j := 0; j < rand.Intn(len(states)); j++ {
			if i != j {
				registered[states[j]] = Valid
			}
		}
		transitions[states[i]] = registered
	}

	return Transitions[T]{transitions}
}
