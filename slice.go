package itermania

import "iter"

// FromSlice creates a generator which iterates over the slice.
func FromSlice[V any](values []V) Gen[V] {
	return func() iter.Seq[V] {
		return func(yield func(V) bool) {
			for _, v := range values {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// ToSlice produces a slice iterated from gen.
//
// Caution: This hangs up if the iteration is infinite.
func ToSlice[V any](gen Gen[V]) []V {
	seq := gen()
	values := []V{}

	for v := range seq {
		values = append(values, v)
	}

	return values
}
