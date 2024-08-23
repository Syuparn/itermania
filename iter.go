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
			for range n {
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

// Dec returns a generator of integers decreasing by one from v.
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

// Range returns a generator of integer range.
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

// Where returns a generator that iterates values only when condGen is true.
func Where[V any](gen Gen[V], condGen Gen[bool]) Gen[V] {
	return func() iter.Seq[V] {
		return func(yield func(V) bool) {
			seq := gen()

			condSeq := condGen()
			condNext, condStop := iter.Pull(condSeq)
			defer condStop()

			for v := range seq {
				cond, ok := condNext()
				if !ok {
					return
				}

				// skip if cond does not meet
				if !cond {
					continue
				}

				if !yield(v) {
					return
				}
			}
		}
	}
}

// Bind applies f to each values iterated from gen.
func Bind[V, W any](gen Gen[V], f func(V) Gen[W]) Gen[W] {
	return func() iter.Seq[W] {
		return func(yield func(W) bool) {
			seq := gen()

			for vVal := range seq {
				wGen := f(vVal)
				wSeq := wGen()

				for wVal := range wSeq {
					if !yield(wVal) {
						return
					}
				}
			}
		}
	}
}

// If works as an if-expression for generators.
//
// NOTE: regardless of cond, both then and else are always evaluated
func If[V any](condGen Gen[bool], thenGen Gen[V], elseGen Gen[V]) Gen[V] {
	return func() iter.Seq[V] {
		return func(yield func(V) bool) {
			condSeq := condGen()
			condNext, condStop := iter.Pull(condSeq)
			defer condStop()

			thenSeq := thenGen()
			thenNext, thenStop := iter.Pull(thenSeq)
			defer thenStop()

			elseSeq := elseGen()
			elseNext, elseStop := iter.Pull(elseSeq)
			defer elseStop()

			for {
				c, ok := condNext()
				if !ok {
					return
				}
				t, ok := thenNext()
				if !ok {
					return
				}
				e, ok := elseNext()
				if !ok {
					return
				}

				result := e
				if c {
					result = t
				}

				if !yield(result) {
					return
				}
			}
		}
	}
}

func All(gen Gen[bool]) Gen[bool] {
	return func() iter.Seq[bool] {
		return func(yield func(bool) bool) {
			seq := gen()
			for v := range seq {
				if !v {
					yield(false)
					return
				}
			}

			yield(true)
		}
	}
}

func Any(gen Gen[bool]) Gen[bool] {
	return func() iter.Seq[bool] {
		return func(yield func(bool) bool) {
			seq := gen()
			for v := range seq {
				if v {
					yield(true)
					return
				}
			}

			yield(false)
		}
	}
}

// Loop returns a generator to iterate the argument v infinitely.
func Loop[V any](v V) Gen[V] {
	return func() iter.Seq[V] {
		return func(yield func(V) bool) {
			for {
				if !yield(v) {
					return
				}
			}
		}
	}
}
