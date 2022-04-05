package queue

import "github.com/flyingdice/proton-pack/internal/validation"

// ErrChannelMustBeSet is the validation check error returned when
// the queue is created with a nil channel.
var ErrChannelMustBeSet = validation.NewError(
	"queue_channel_must_be_set",
	"the queue channel must be set and cannot be an nil",
)

func defaultChecks[T any]() []validation.Check[*Queue[T]] {
	return []validation.Check[*Queue[T]]{
		checkChannelSet[T](),
	}
}

// checkChannelSet validates queue channel is set..
func checkChannelSet[T any]() validation.Check[*Queue[T]] {
	return func(q *Queue[T]) *validation.Error {
		if q.ch == nil {
			return ErrChannelMustBeSet
		}
		return nil
	}
}
