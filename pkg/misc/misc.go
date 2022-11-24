package misc

// Must is a shorthand to discard the error return value usually returned by
// functions, which it achieves by accepting two arguments (the second of which
// must be an error), and returning the first one, of any value.
func Must[T any](a T, _ error) T {
	return a
}
