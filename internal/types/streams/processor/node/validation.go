package node

import (
	"github.com/flyingdice/proton-pack/internal/validation"
)

// ErrChildrenByNameNotSet is the validation check error returned when
// the topic doesn't match length or character requirements.
var ErrChildrenByNameNotSet = validation.NewError(
	"node_children_by_name_not_set",
	"the childrenByName map must be allocated and not nil",
)

var defaultChecks = []validation.Check[*Node]{
	checkChildrenByNameIsSet(),
}

// checkChildrenByNameIsSet validates childrenByName map has been allocated.
func checkChildrenByNameIsSet() validation.Check[*Node] {
	return func(n *Node) *validation.Error {
		if n.childrenByName == nil {
			return ErrChildrenByNameNotSet
		}
		return nil
	}
}
