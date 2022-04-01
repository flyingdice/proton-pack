package timestamp

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/validation"
	"time"
)

// ErrMustBeNewerThanUnixEpoch is the validation check error returned when
// the timestamp is before the unix epoch (Jan 1st, 1970).
var ErrMustBeNewerThanUnixEpoch = validation.NewError(
	"timestamp_must_be_newer_than_unix_epoch",
	fmt.Sprintf("the timestamp value must be newer than %s", unixEpoch),
)

var unixEpoch = time.Unix(0, 0)

var defaultChecks = []validation.Check[Timestamp]{
	checkNewerThanUnixEpoch(),
}

// checkNewerThanUnixEpoch validates timestamp is newer than unix epoch (Jan 1st, 1970).
func checkNewerThanUnixEpoch() validation.Check[Timestamp] {
	return func(ts Timestamp) *validation.Error {
		if ts.Unix() < unixEpoch.Unix() {
			return ErrMustBeNewerThanUnixEpoch
		}
		return nil
	}
}
