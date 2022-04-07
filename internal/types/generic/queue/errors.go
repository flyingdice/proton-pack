package queue

import "github.com/pkg/errors"

var ErrPushUnbuffered = errors.New("push failed; unbuffered channel")
var ErrEmpty = errors.New("queue is empty")
