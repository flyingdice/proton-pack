package validation

// Checker represents types that can be checked for validation errors.
type Checker interface {
	Check() ErrorGroup
}

// Check is a function that performs a validation check of a value of type T.
type Check[T any] func(val T) *Error

// RunChecks runs checks for the given value.
//
// If the value is invalid, one or more Errors will be returned.
func RunChecks[T any](val T, checks ...Check[T]) ErrorGroup {
	errs := &Errors{}
	for _, check := range checks {
		if e := check(val); e != nil {
			errs.Append(e)
		}
	}
	return errs.NilWhenEmpty()
}

// RunCheckers runs checkers and bundles all errors into a group.
func RunCheckers(checkers ...Checker) ErrorGroup {
	errs := &Errors{}
	for _, checker := range checkers {
		if err := checker.Check(); err != nil {
			errs.Append(err.Errors()...)
		}
	}
	return errs.NilWhenEmpty()
}
