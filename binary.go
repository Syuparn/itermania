package itermania

import "iter"

func And[V bool](xGen Gen[bool], yGen Gen[bool]) Gen[bool] {
	bin := Bin(func(xVal bool, yVal bool) bool {
		return xVal && yVal
	})
	return bin(xGen, yGen)
}

func Or[V bool](xGen Gen[bool], yGen Gen[bool]) Gen[bool] {
	bin := Bin(func(xVal bool, yVal bool) bool {
		return xVal || yVal
	})
	return bin(xGen, yGen)
}

// Bin returns a generator from two generators and a binary operation
func Bin[V, W any](op func(V, V) W) func(Gen[V], Gen[V]) Gen[W] {
	return func(xGen Gen[V], yGen Gen[V]) Gen[W] {
		return func() iter.Seq[W] {
			return func(yield func(W) bool) {
				xSeq := xGen()
				for x := range xSeq {
					ySeq := yGen()

					for y := range ySeq {
						if !yield(op(x, y)) {
							return
						}
					}
				}
			}
		}
	}
}
