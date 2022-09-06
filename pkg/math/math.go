package math

import "golang.org/x/exp/constraints"

// MultipleMin returns the minimum of the values passed as parameters. Starting with 1.19, Go has
// a math.Min() function, which should be favored over this for production, as it takes care of
// more cases such as infinities. However, it only returns the minimum of 2 arguments, so this
// function is to be used for more than 2, where the extra precautions aren't necessary.
func MultipleMin[T constraints.Ordered](first T, more ...T) T {
	for _, n := range more {
		if n < first {
			first = n
		}
	}
	return first
}

// MultipleMax returns the maiximum of the values passed. Same warnings as for MultipleMin apply
// here.
func MultipleMax[T constraints.Ordered](first T, more ...T) T {
	for _, n := range more {
		if n > first {
			first = n
		}
	}
	return first
}
