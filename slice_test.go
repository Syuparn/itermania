package itermania

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	gen := FromSlice([]int{1, 2})
	seq := gen()
	next, stop := iter.Pull(seq)
	defer stop()

	v, ok := next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)

	v, ok = next()
	assert.False(t, ok)
}

func TestToSlice(t *testing.T) {
	tests := []struct {
		name     string
		gen      Gen[int]
		expected []int
	}{
		{
			"one element",
			Const(1),
			[]int{1},
		},
		{
			"three elements",
			Range(1, 4, 1),
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToSlice(tt.gen)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
