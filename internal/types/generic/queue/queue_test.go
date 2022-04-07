package queue

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/testing/assertion"
	"testing"
	"testing/quick"
)

// TestQueue_PushPop checks Pop/Push interact as expected.
func TestQueue_PushPop(t *testing.T) {
	assert := assertion.Fatal(t)

	q := &Queue[string]{ch: make(chan string, 1)}

	// Pop from empty queue is ErrEmpty.
	_, err := q.Pop()
	assert.Equal(err, ErrEmpty)

	// Push new value without problems.
	err = q.Push("foo")
	assert.OK(err)

	// Pop value back off successfully.
	item, err := q.Pop()
	assert.OK(err)
	assert.Equal(item, "foo")
}

// TestQueue_PushUnbuffered checks error is returned if push to a channel that is not buffered.
func TestQueue_PushUnbuffered(t *testing.T) {
	assert := assertion.Fatal(t)

	q := &Queue[string]{ch: make(chan string)}

	err := q.Push("foo")
	assert.Equal(err, ErrPushUnbuffered)
}

// TestQueue_String checks String() output is expected format.
func TestQueue_String(t *testing.T) {
	assert := assertion.Error(t)
	checker := func() bool {
		q := &Queue[string]{ch: make(chan string, 10)}
		got := q.String()
		want := fmt.Sprintf("Queue[%T](len=%d cap=%d)", *new(string), q.Len(), q.Cap())
		return assert.Equal(got, want)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
