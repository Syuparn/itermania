package itermania

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// Gen is a generator function that generates a new iterator
type Gen[V any] = func() iter.Seq[V]

type Number interface {
	constraints.Integer | constraints.Float
}
