package validation

import (
	"errors"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"
	"testing/quick"
)

// TestCheckError_Error checks Error() output is expected format.
func Test_Validate(t *testing.T) {
	var errTestIntPositive = NewError("int_must_be_positive", "integer value must be positive")

	// Simple check function that tests integers are positive.
	checkPositive := func() Check[int] {
		return func(x int) *Error {
			if x%2 != 0 {
				return errTestIntPositive
			}
			return nil
		}
	}

	checker := func(v int) bool {
		err := Validate(v, checkPositive())
		if v%2 == 0 {
			if err != nil {
				t.Errorf("got err %s; expected none", err)
				return false
			}
		} else {
			if !errors.Is(err, errTestIntPositive) {
				t.Errorf("got err %s; expected %s", err, errTestIntPositive)
				return false
			}
		}
		return true
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}

// TestCheckError_Error checks Error() output is expected format.
func TestCheckError_Error(t *testing.T) {
	checker := func() bool {
		slug, desc := faker.Word(), faker.Sentence()
		ce := NewError(slug, desc)
		return ce.Error() == fmt.Sprintf(
			"Error(code=%d slug=%s desc=%s)",
			ce.Code,
			ce.Slug,
			ce.Desc,
		)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error()
	}
}

// TestCheckError_Code sets the code attribute is set upon creation to hash code of the slug.
func TestCheckError_Code(t *testing.T) {
	checker := func() bool {
		slug, desc := faker.Word(), faker.Sentence()
		ce := NewError(slug, desc)
		return ce.Code == errCode(slug)
	}
	if err := quick.Check(checker, nil); err != nil {
		t.Error(err)
	}
}
