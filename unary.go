package itermania

import "iter"

func Not(gen Gen[bool]) Gen[bool] {
	uni := Uni(func(v bool) bool {
		return !v
	})
	return uni(gen)
}

// Uni returns a generator from a generators and a unary operation
func Uni[V, W any](op func(V) W) func(Gen[V]) Gen[W] {
	return func(xGen Gen[V]) Gen[W] {
		return func() iter.Seq[W] {
			return func(yield func(W) bool) {
				xSeq := xGen()
				for x := range xSeq {
					if !yield(op(x)) {
						return
					}
				}
			}
		}
	}
}
