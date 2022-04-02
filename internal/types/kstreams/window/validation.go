package window

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/validation"
	"time"
)

// ErrLoMustBeNewerThanUnixEpoch is the validation check error returned when
// the window is before the unix epoch (Jan 1st, 1970).
var ErrLoMustBeNewerThanUnixEpoch = validation.NewError(
	"window_lo_must_be_newer_than_unix_epoch",
	fmt.Sprintf("the window lo value must be newer than %s", unixEpoch),
)

// ErrHiGreaterThanLo is the validation check error returned when
// the window hi is before the lo.
var ErrHiGreaterThanLo = validation.NewError(
	"window_hi_must_be_greater_than_lo",
	"the window hi value must be greater than the lo value",
)

var unixEpoch = time.Unix(0, 0)

var defaultChecks = []validation.Check[Window]{
	checkLoNewerThanUnixEpoch(),
	checkHiGreaterThanLo(),
}

// checkLoNewerThanUnixEpoch validates window lo is newer than unix epoch (Jan 1st, 1970).
func checkLoNewerThanUnixEpoch() validation.Check[Window] {
	return func(w Window) *validation.Error {
		if w.lo.Unix() < unixEpoch.Unix() {
			return ErrLoMustBeNewerThanUnixEpoch
		}
		return nil
	}
}

// checkHiGreaterThanLo validates window hi is greater than lo.
func checkHiGreaterThanLo() validation.Check[Window] {
	return func(w Window) *validation.Error {
		if w.hi.Before(w.lo) {
			return ErrHiGreaterThanLo
		}
		return nil
	}
}
