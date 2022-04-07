package collector

import (
	"github.com/flyingdice/proton-pack/internal/testing/assertion"
	"testing"
)

// TestState performs a sanity check of state constants/configuration.
//
// This exists mainly to bubble up misconfiguration or logic changes and
// require the developer to explicitly make them known.
func TestState(t *testing.T) {
	assert := assertion.Fatal(t)

	// Initial state is Closed.
	assert.Equal(Initial, Closed)

	// Registered states are expected.
	assert.Equal(States[0], Opened)
	assert.Equal(States[1], Closed)

	// Opened can transition to closed and not itself.
	opened, ok := Transitions[Opened]
	assert.Equal(ok, true)
	assert.Equal(opened[Closed], true)
	assert.Equal(opened[Opened], false)

	// Closed can transition to opened and not itself.
	closed, ok := Transitions[Closed]
	assert.Equal(ok, true)
	assert.Equal(closed[Opened], true)
	assert.Equal(closed[Closed], false)
}
