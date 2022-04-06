package store

import (
	"github.com/matryer/is"
	"testing"
)

// TestValidation_MapNotNil checks that validation inspects the underlying map
// to make sure its allocated.
func TestValidation_MapNotNil(t *testing.T) {
	assert := is.New(t)

	_, err := NewStore()
	assert.NoErr(err)

	s := &Store{}
	errs := s.Check()

	assert.Equal(errs.Error(), ErrStoreMapIsNil.Error())
}
