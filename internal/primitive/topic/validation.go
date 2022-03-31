package topic

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/validation"
	"regexp"
)

var pattern = `[a-zA-Z0-9_.\-]{1,255}`

var defaultChecks = []validation.Check[Topic]{
	checkPattern(regexp.MustCompile(pattern)),
}

var ErrMustMatchPattern = validation.NewCheckError(
	"topic_must_match_pattern",
	fmt.Sprintf("the topic value must match the following pattern: %s", pattern),
)

func checkPattern(r *regexp.Regexp) validation.Check[Topic] {
	return func(t Topic) *validation.CheckError {
		match := r.MatchString(string(t))
		if !match {
			return ErrMustMatchPattern
		}
		return nil
	}
}
