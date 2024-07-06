package itermania

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	t.Run("const int", func(t *testing.T) {
		gen := Const(5)
		seq := gen()

		next, stop := iter.Pull(seq)
		defer stop()

		v, ok := next()
		assert.Equal(t, 5, v)
		assert.True(t, ok)

		v, ok = next()
		assert.False(t, ok)
	})

	t.Run("const string", func(t *testing.T) {
		gen := Const("foo")
		seq := gen()

		next, stop := iter.Pull(seq)
		defer stop()

		v, ok := next()
		assert.Equal(t, "foo", v)
		assert.True(t, ok)

		v, ok = next()
		assert.False(t, ok)
	})
}

func TestHead(t *testing.T) {
	tests := []struct {
		name     string
		gen      Gen[int]
		n        int
		expected []int
	}{
		{
			"empty",
			Inc(0),
			0,
			[]int{},
		},
		{
			"first 3 elems",
			Inc(0),
			3,
			[]int{0, 1, 2},
		},
		{
			"more than iterator elems",
			Const(10),
			3,
			[]int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Head(tt.gen, tt.n)

			actual := []int{}
			for e := range gen() {
				actual = append(actual, e)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestInc(t *testing.T) {
	gen := Inc(10)
	seq := gen()

	next, stop := iter.Pull(seq)
	defer stop()

	v, ok := next()
	assert.Equal(t, 10, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 11, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 12, v)
	assert.True(t, ok)
}

func TestDec(t *testing.T) {
	gen := Dec(10)
	seq := gen()

	next, stop := iter.Pull(seq)
	defer stop()

	v, ok := next()
	assert.Equal(t, 10, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 9, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 8, v)
	assert.True(t, ok)
}

func TestRange(t *testing.T) {
	tests := []struct {
		name     string
		gen      Gen[int]
		expected []int
	}{
		{
			"1 to 10",
			Range(1, 11, 1),
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			"1 to 10 by 2",
			Range(1, 11, 2),
			[]int{1, 3, 5, 7, 9},
		},
		{
			"10 to 0 by -2",
			Range(10, -1, -2),
			[]int{10, 8, 6, 4, 2, 0},
		},
		{
			"20 to 10 by 1",
			Range(20, 10, 1),
			[]int{},
		},
		{
			"10 to 10 by 1",
			Range(10, 10, 1),
			[]int{},
		},
		{
			"10 to 20 by -1",
			Range(10, 20, -1),
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := []int{}
			for e := range tt.gen() {
				actual = append(actual, e)
			}

			assert.Equal(t, tt.expected, actual)
		})
	}
}
