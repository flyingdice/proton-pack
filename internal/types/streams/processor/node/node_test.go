package node

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/testing/assertion"
	"testing"
	"testing/quick"
)

// TestNode_Name checks Name() output is expected format.
func TestNode_Name(t *testing.T) {
	assert := assertion.Error(t)
	checker := func(n *Node) bool {
		return assert.Equal(n.Name(), n.name)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestNode_String checks String() output is expected format.
func TestNode_String(t *testing.T) {
	assert := assertion.Error(t)
	checker := func(n *Node) bool {
		return assert.Equal(n.String(), fmt.Sprintf("Node(name=%s)", n.name))
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestNode_AddChild checks adding children is tracked as expected.
func TestNode_AddChild(t *testing.T) {
	assert := assertion.Error(t)
	checker := func(n *Node, c *Node) bool {
		x, y := len(n.children), len(n.childrenByName)

		// Assert error returned on name collision.
		if _, found := n.childrenByName[c.Name()]; found {
			err := n.AddChild(c)
			return assert.Equal(err.Error(), fmt.Sprintf("child '%s' already added", c.Name()))
		} else {
			if !assert.OK(n.AddChild(c)) {
				return false
			}
		}

		// Assert collection lengths increased by 1.
		return assert.Equal(len(n.children), x+1) && assert.Equal(len(n.childrenByName), y+1)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestNode_GetChild checks fetching children works as expected.
func TestNode_GetChild(t *testing.T) {
	assert := assertion.Error(t)
	checker := func(n *Node) bool {
		want := n.children[0]

		c := n.GetChild(want.Name())
		if !assert.Equal(c.Name(), want.Name()) {
			return false
		}

		c = n.GetChild(want.Name() + "foobar")
		return assert.Equal(c, (*Node)(nil))
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
