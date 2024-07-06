package itermania

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// Const returns a generator to iterate the argument v once.
func Const[V any](v V) Gen[V] {
	return func() iter.Seq[V] {
		return func(yield func(V) bool) {
			if !yield(v) {
				return
			}
		}
	}
}

// Head returns a generator to iterator first n elements in gen.
func Head[V any](gen Gen[V], n int) func() iter.Seq[V] {
	return func() iter.Seq[V] {
		return func(yield func(V) bool) {
			seq := gen()
			next, stop := iter.Pull(seq)
			defer stop()
			for _ = range n {
				v, ok := next()
				if !ok {
					return
				}
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Inc returns a generator of integers increasing by one from v.
func Inc[V constraints.Integer](v V) Gen[V] {
	return func() iter.Seq[V] {
		return func(yield func(V) bool) {
			for {
				if !yield(v) {
					return
				}
				v++
			}
		}
	}
}

// Inc returns a generator of integers decreasing by one from v.
func Dec[V constraints.Integer](v V) Gen[V] {
	return func() iter.Seq[V] {
		return func(yield func(V) bool) {
			for {
				if !yield(v) {
					return
				}
				v--
			}
		}
	}
}

// Inc returns a generator of integer range.
func Range[V constraints.Integer](start, stop, step V) Gen[V] {
	return func() iter.Seq[V] {
		return func(yield func(V) bool) {
			i := start
			increasing := step > 0
			for {
				if (increasing && (i >= stop)) || (!increasing && (i <= stop)) {
					return
				}

				if !yield(i) {
					return
				}
				i += step
			}
		}
	}
}
