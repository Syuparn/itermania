package itermania

import (
	"iter"

	"golang.org/x/exp/constraints"
)

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

func Eq[V comparable](xGen Gen[V], yGen Gen[V]) Gen[bool] {
	bin := Bin(func(xVal V, yVal V) bool {
		return xVal == yVal
	})
	return bin(xGen, yGen)
}

func Neq[V comparable](xGen Gen[V], yGen Gen[V]) Gen[bool] {
	bin := Bin(func(xVal V, yVal V) bool {
		return xVal != yVal
	})
	return bin(xGen, yGen)
}

func Gt[V constraints.Ordered](xGen Gen[V], yGen Gen[V]) Gen[bool] {
	bin := Bin(func(xVal V, yVal V) bool {
		return xVal > yVal
	})
	return bin(xGen, yGen)
}

func Lt[V constraints.Ordered](xGen Gen[V], yGen Gen[V]) Gen[bool] {
	bin := Bin(func(xVal V, yVal V) bool {
		return xVal < yVal
	})
	return bin(xGen, yGen)
}

func Ge[V constraints.Ordered](xGen Gen[V], yGen Gen[V]) Gen[bool] {
	bin := Bin(func(xVal V, yVal V) bool {
		return xVal >= yVal
	})
	return bin(xGen, yGen)
}

func Le[V constraints.Ordered](xGen Gen[V], yGen Gen[V]) Gen[bool] {
	bin := Bin(func(xVal V, yVal V) bool {
		return xVal <= yVal
	})
	return bin(xGen, yGen)
}

func Add[V constraints.Ordered](xGen Gen[V], yGen Gen[V]) Gen[V] {
	bin := Bin(func(xVal V, yVal V) V {
		return xVal + yVal
	})
	return bin(xGen, yGen)
}

func Sub[V Number](xGen Gen[V], yGen Gen[V]) Gen[V] {
	bin := Bin(func(xVal V, yVal V) V {
		return xVal - yVal
	})
	return bin(xGen, yGen)
}

func Mul[V Number](xGen Gen[V], yGen Gen[V]) Gen[V] {
	bin := Bin(func(xVal V, yVal V) V {
		return xVal * yVal
	})
	return bin(xGen, yGen)
}

func Div[V Number](xGen Gen[V], yGen Gen[V]) Gen[V] {
	bin := Bin(func(xVal V, yVal V) V {
		return xVal / yVal
	})
	return bin(xGen, yGen)
}

func Mod[V constraints.Integer](xGen Gen[V], yGen Gen[V]) Gen[V] {
	bin := Bin(func(xVal V, yVal V) V {
		return xVal % yVal
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
