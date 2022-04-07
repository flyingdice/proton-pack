package store

import (
	"github.com/flyingdice/proton-pack/internal/testing/assertion"
	"testing"
)

// TestValidation_MapNotNil checks that validation inspects the underlying map
// to make sure its allocated.
func TestValidation_MapNotNil(t *testing.T) {
	assert := assertion.Fatal(t)

	_, err := New()
	assert.OK(err)

	s := &Store{}
	errs := s.Check()

	assert.Equal(errs, ErrStoreMapIsNil)
}
