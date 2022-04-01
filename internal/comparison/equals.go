package comparison

// Equaler is implemented by any value that has a Equals method,
// which defines the equality comparison for that value.
type Equaler interface {
	Equals(v any) bool
}
