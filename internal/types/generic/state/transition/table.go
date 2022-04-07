package transition

import (
	"github.com/bxcodec/faker/v3"
	"github.com/flyingdice/proton-pack/internal/types/generic/state/state"
	"github.com/flyingdice/proton-pack/internal/validation"
	"math/rand"
	"reflect"
)

// Table represents a nested map for storing states and their valid transitions.
type Table[T state.State] map[T]map[T]bool

// Transitions represents a structure for describing states and the valid transitions between them.
type Transitions[T state.State] struct {
	table Table[T]
}

// New creates and validates a new Transitions.
func New[T state.State](transitions Table[T]) (*Transitions[T], validation.ErrorGroup) {
	t := &Transitions[T]{transitions}
	return t, t.Check()
}

// Check runs default validation checks for the Transitions.
func (t *Transitions[T]) Check() validation.ErrorGroup {
	return validation.RunChecks[*Transitions[T]](t, defaultChecks[T]()...)
}

// Valid return true for registered transition between current -> next states.
func (t *Transitions[T]) Valid(current, next T) bool {
	c, ok := t.table[current]
	if !ok {
		return false
	}
	n, ok := c[next]
	if !ok {
		return false
	}
	return n
}

// Len returns number of registered transitions for all states.
func (t *Transitions[T]) Len() int {
	return len(t.table)
}

// Generate random Transition values.
//
// Interface: quick.Generator
func (Transitions[T]) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate[T](rand, []T{}))
}

// Generate from a set of random states.
func Generate[T state.State](rand *rand.Rand, states []T) *Transitions[T] {
	faker.SetRandomSource(rand)

	if len(states) == 0 {
		states = state.Generate[T](rand)
	}

	// Register random transitions for states within the given set.
	transitions := make(map[T]map[T]bool)
	for i := 0; i < rand.Intn(len(states)); i++ {
		registered := make(map[T]bool)
		for j := 0; j < rand.Intn(len(states)); j++ {
			if i != j {
				registered[states[j]] = true
			}
		}
		transitions[states[i]] = registered
	}

	return &Transitions[T]{transitions}
}
