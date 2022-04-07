package node

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/flyingdice/proton-pack/internal/types/generic/state/machine"
	"github.com/flyingdice/proton-pack/internal/types/streams/processor"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/context"
	"github.com/flyingdice/proton-pack/internal/types/streams/record/record"
	"github.com/flyingdice/proton-pack/internal/validation"
	"github.com/pkg/errors"
	"math/rand"
	"reflect"
	"testing/quick"
)

var _ fmt.Stringer = (*Node)(nil)
var _ quick.Generator = (*Node)(nil)

// Node represents a node within a topology graph.
type Node struct {
	children       []*Node
	childrenByName map[string]*Node
	processor      processor.Processor
	name           string
	machine        *machine.Machine[State]
}

// New creates and validates a new Node from the given values.
func New(name string, processor_ processor.Processor) (*Node, validation.ErrorGroup) {
	m, err := machine.New[State](Initial, States, Transitions)
	if err != nil {
		return nil, err
	}

	n := &Node{
		childrenByName: make(map[string]*Node),
		name:           name,
		processor:      processor_,
		machine:        m,
	}
	return n, n.Check()
}

// Check runs default validation checks for the node.
func (n *Node) Check() validation.ErrorGroup {
	return validation.RunChecks[*Node](n, defaultChecks...)
}

// Name of the node.
//
// Interface: Node
func (n *Node) Name() string { return n.name }

// AddChild adds a child to this node.
//
// Interface: Node
func (n *Node) AddChild(c *Node) error {
	if _, found := n.childrenByName[c.Name()]; found {
		return errors.Errorf("child '%s' already added", c.Name())
	}
	n.children = append(n.children, c)
	n.childrenByName[c.Name()] = c
	return nil
}

// GetChild fetches a child node by its name.
//
// Interface: Node
func (n *Node) GetChild(name string) *Node {
	return n.childrenByName[name]
}

// Children returns slice of all registered children for this node.
//
// Interface: Node
func (n *Node) Children() []*Node {
	return n.children
}

// Open the node.
//
// The node can be re-used by calling Close and Open again.
//
// Interface: Processor
func (n *Node) Open(ctx processor.Context) error {
	return n.machine.To(Opened, func() error {
		return n.processor.Open(ctx)
	})
}

// Process a single record.
//
// It is required that Open be called before any records are processed.
// It is required that Close not be called with records still to process.
//
// Interface: Processor
func (n *Node) Process(ctx context.Context, record record.Record) error {
	return n.machine.MustBe(Opened, func() error {
		return n.processor.Process(ctx, record)
	})
}

// Close the node.
//
// It can be re-used by calling Open again.
//
// Interface: Processor
func (n *Node) Close() error {
	return n.machine.To(Closed, func() error {
		return n.processor.Close()
	})
}

// String value of the node.
//
// Interface: fmt.Stringer.
func (n *Node) String() string {
	return fmt.Sprintf("Node(name=%s)", n.name)
}

// Generate random node values.
//
// Interface: quick.Generator
func (*Node) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Generate(rand, rand.Intn(10-1)+1))
}

// Generate random node value.
func Generate(rand *rand.Rand, numChildren int) *Node {
	faker.SetRandomSource(rand)

	sm, err := machine.New[State](Initial, States, Transitions)
	if err != nil {
		panic(err)
	}

	children := make([]*Node, 0, numChildren)
	childrenByName := make(map[string]*Node, numChildren)

	for i := 0; i < numChildren; i++ {
		child := Generate(rand, rand.Intn((numChildren/2)+1))
		if _, ok := childrenByName[child.name]; !ok {
			children = append(children, child)
			childrenByName[child.name] = child
		}
	}

	return &Node{
		machine:        sm,
		children:       children,
		childrenByName: childrenByName,
		name:           faker.Word(),
		processor:      nil,
	}
}
