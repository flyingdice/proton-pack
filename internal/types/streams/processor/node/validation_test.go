package node

import (
	"github.com/flyingdice/proton-pack/internal/testing/assertion"
	"testing"
	"testing/quick"
)

// TestValidation_NewNode checks that default validation checks are run.
func TestValidation_NewNode(t *testing.T) {
	assert := assertion.Error(t)
	checker := func(name string) bool {
		_, err := New(name, nil)
		return assert.OK(err)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestValidation_ChildrenByNameNotSet checks that childrenByName is allocated.
func TestValidation_ChildrenByNameNotSet(t *testing.T) {
	assert := assertion.Error(t)
	checker := func() bool {
		n := &Node{}
		err := n.Check()
		if !assert.NotOK(err) {
			return false
		}
		return assert.Equal(err, ErrChildrenByNameNotSet)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
