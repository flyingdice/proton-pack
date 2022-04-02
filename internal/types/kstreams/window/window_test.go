package window

import (
	"fmt"
	"testing"
	"testing/quick"
	"time"
)

// TestWindow_String checks String() output is expected format.
func TestWindow_String(t *testing.T) {
	checker := func(w Window) bool {
		return w.String() == fmt.Sprintf(
			"Window(lo=%s hi=%s)",
			w.lo.Format(time.RFC3339),
			w.hi.Format(time.RFC3339),
		)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestWindow_Duration checks Duration() output is expected.
func TestWindow_Duration(t *testing.T) {
	checker := func(w Window) bool {
		return w.Duration() == w.hi.Sub(w.lo)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestWindow_Overlaps checks time overlap between two Window instances.
func TestWindow_Overlaps(t *testing.T) {
	checker := func(w1 Window, ns int64) bool {
		if !w1.Overlaps(w1) {
			t.Error("window should always overlap itself")
			return false
		}
		w2 := Window{
			lo: w1.lo.Add(time.Duration(ns)),
			hi: w1.hi.Add(time.Duration(ns)),
		}
		overlaps := w1.Overlaps(w2)
		if w1.lo.Before(w2.hi) && w2.lo.Before(w1.hi) {
			return overlaps
		} else {
			return !overlaps
		}
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestWindow_EqualsTrue checks equality between two Window instances and their
// time.Time representations.
func TestWindow_EqualsTrue(t *testing.T) {
	checker := func(w Window) bool {
		return w.Equals(w)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestWindow_EqualsFalse checks inequality between two in-equal Window instances and their
// time.Time representations.
func TestWindow_EqualsFalse(t *testing.T) {
	checker := func(w1 Window) bool {
		w2 := Window{
			lo: w1.lo.Add(1 * time.Second),
			hi: w1.hi.Add(1 * time.Second),
		}
		return !w1.Equals(w2)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
