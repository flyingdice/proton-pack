package machine

import (
	"github.com/matryer/is"
	"testing"
)

// TestErrAlreadyInState_Error checks Error() output for ErrAlreadyInState.
func TestErrAlreadyInState_Error(t *testing.T) {
	assert := is.New(t)

	err := ErrAlreadyInState[string]{State: "foo"}
	assert.Equal(err.Error(), "machine already in state 'foo'")
}

// TestErrNotInState_Error checks Error() output for ErrNotInState.
func TestErrNotInState_Error(t *testing.T) {
	assert := is.New(t)

	err := ErrNotInState[string]{State: "foo"}
	assert.Equal(err.Error(), "machine not in state 'foo'")
}

// TestErrInvalidTransition_Error checks Error() output for ErrInvalidTransition.
func TestErrInvalidTransition_Error(t *testing.T) {
	assert := is.New(t)

	err := ErrInvalidTransition[string]{Current: "foo", Next: "bar"}
	assert.Equal(err.Error(), "machine cannot transition from 'foo' to 'bar'")
}
