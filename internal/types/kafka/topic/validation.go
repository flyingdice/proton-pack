package topic

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/validation"
	"regexp"
)

// ErrMustMatchPattern is the validation check error returned when
// the topic doesn't match length or character requirements.
var ErrMustMatchPattern = validation.NewError(
	"topic_must_match_pattern",
	fmt.Sprintf("the topic value must match the following pattern: %s", pattern),
)

var pattern = `[a-zA-Z0-9_.\-]{1,255}`

var defaultChecks = []validation.Check[Topic]{
	checkPattern(regexp.MustCompile(pattern)),
}

// checkPattern validates topic matches a specific regex pattern to meet length/character requirements.
func checkPattern(r *regexp.Regexp) validation.Check[Topic] {
	return func(t Topic) *validation.Error {
		match := r.MatchString(string(t))
		if !match {
			return ErrMustMatchPattern
		}
		return nil
	}
}
