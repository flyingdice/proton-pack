package collector

import (
	"github.com/matryer/is"
	"testing"
)

// TestState performs a sanity check of state constants/configuration.
//
// This exists mainly to bubble up misconfiguration or logic changes and
// require the developer to explicitly make them known.
func TestState(t *testing.T) {
	assert := is.New(t)

	// Initial state is Closed.
	assert.Equal(Initial, Closed)

	// Registered states are expected.
	assert.Equal(States[0], Opened)
	assert.Equal(States[1], Closed)

	// Opened can transition to closed and not itself.
	opened, ok := Transitions[Opened]
	assert.True(ok)
	assert.True(opened[Closed])
	assert.True(!opened[Opened])

	// Closed can transition to opened and not itself.
	closed, ok := Transitions[Closed]
	assert.True(ok)
	assert.True(closed[Opened])
	assert.True(!closed[Closed])
}
