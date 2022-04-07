package queue

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/validation"
)

// Queue represents a FIFO queue.
type Queue[T any] struct {
	ch chan T
}

// NewQueue creates and validates a new Queue.
func NewQueue[T any](cap int) (*Queue[T], validation.ErrorGroup) {
	q := &Queue[T]{
		ch: make(chan T, cap),
	}
	return q, q.Check()
}

// Check runs default validation checks for the Queue.
func (q *Queue[T]) Check() validation.ErrorGroup {
	return validation.RunChecks[*Queue[T]](q, defaultChecks[T]()...)
}

// Len returns length of queue.
func (q *Queue[T]) Len() int { return len(q.ch) }

// Cap returns capacity of queue.
func (q *Queue[T]) Cap() int { return cap(q.ch) }

// Push places given value in the queue.
func (q *Queue[T]) Push(val T) error {
	select {
	case q.ch <- val:
		return nil
	default:
		return ErrPushUnbuffered
	}
}

// Pop removes head value in the queue.
func (q *Queue[T]) Pop() (T, error) {
	select {
	case val := <-q.ch:
		return val, nil
	default:
		return *new(T), ErrEmpty
	}
}

// String value of the Queue.
//
// Interface: fmt.Stringer.
func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue[%T](len=%d cap=%d)", *new(T), q.Len(), q.Cap())
}
