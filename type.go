package itermania

import "iter"

// Gen is a generator function that generates a new iterator
type Gen[V any] = func() iter.Seq[V]
