package assertion

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
)

// defaultOptions for go-cmp comparisons.
var defaultOptions = cmp.Options{
	cmpopts.EquateErrors(),
	cmpopts.EquateEmpty(),
}

// log is the func called when a boolean condition is not met.
type log func(format string, args ...any)

type Assert struct {
	tb   testing.TB
	log  log
	opts cmp.Options
}

// Error asserts using Error logger.
//
// Helpful for consuming the boolean result of an assertion.
func Error(tb testing.TB, opts ...cmp.Option) Assert {
	return New(tb, tb.Errorf, opts...)
}

// Fatal asserts using Fatal logger.
//
// Helpful for immediately aborting test execution on failed assertion.
func Fatal(tb testing.TB, opts ...cmp.Option) Assert {
	return New(tb, tb.Fatalf, opts...)
}

// New creates a new Assert with the given params.
func New(tb testing.TB, log log, opts ...cmp.Option) Assert {
	opts = append(opts, defaultOptions...)
	return Assert{tb, log, opts}
}

// Assert fails the test if the condition is false.
func (a Assert) Assert(condition bool, format string, args ...any) bool {
	a.tb.Helper()
	if !condition {
		a.log(format, args...)
		return false
	}
	return true
}

// OK fails the test if err is not nil.
func (a Assert) OK(err error) bool {
	a.tb.Helper()
	if err != nil {
		a.log("unexpected error: %s", err.Error())
		return false
	}
	return true
}

// NotOK fails the test if err is nil.
func (a Assert) NotOK(err error) bool {
	a.tb.Helper()
	if err == nil {
		a.log("expected error; got none")
		return false
	}
	return true
}

// Equal fails the test if got is not equal to want.
func (a Assert) Equal(got, want any) bool {
	a.tb.Helper()
	if !cmp.Equal(got, want, a.opts) {
		diff := cmp.Diff(got, want, a.opts)
		a.log("\n\tgot: %#v\n\n\twant: %#v\n\n\tdiff: %s", got, want, diff)
		return false
	}
	return true
}
