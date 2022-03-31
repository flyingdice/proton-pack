package timestamp

import (
	"fmt"
	"github.com/flyingdice/proton-pack/internal/validation"
	"time"
)

var unixEpoch = time.Unix(0, 0)

var defaultChecks = []validation.Check[Timestamp]{
	checkNewerThanUnixEpoch(),
}

var ErrMustBeNewerThanUnixEpoch = validation.NewCheckError(
	"timestamp_must_be_newer_than_unix_epoch",
	fmt.Sprintf("the timestamp value must be newer than %s", unixEpoch),
)

func checkNewerThanUnixEpoch() validation.Check[Timestamp] {
	return func(ts Timestamp) validation.CheckError {
		if ts.Unix() < unixEpoch.Unix() {
			return ErrMustBeNewerThanUnixEpoch
		}
		return nil
	}
}
