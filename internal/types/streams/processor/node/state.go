package node

// State represents the valid states of a node.
//
// Interface: state.State
type State string

const (
	Opened  State = "opened"
	Closed  State = "closed"
	Initial       = Closed
)

// States is list of all valid node states.
var States = []State{
	Opened,
	Closed,
}

// Transitions contains valid transitions between states of a node.
var Transitions = map[State]map[State]bool{
	Opened: {Closed: true},
	Closed: {Opened: true},
}
