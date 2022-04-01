package validation

// Checker represents types that can be checked for validation errors.
type Checker interface {
	Check() ErrorGroup
}

// Check is a function that performs a validation check of a value of type T.
type Check[T any] func(val T) *Error

// Validate runs checks for the given value.
//
// If the value is invalid, one or more Errors will be returned.
func Validate[T any](val T, checks ...Check[T]) ErrorGroup {
	errs := &Errors{}

	for _, check := range checks {
		if e := check(val); e != nil {
			errs.Append(e)
		}
	}

	return errs.NilWhenEmpty()
}
