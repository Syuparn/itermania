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

func TestWhere(t *testing.T) {
	tests := []struct {
		name     string
		gen      Gen[int]
		cond     Gen[bool]
		expected []int
	}{
		{
			"only true",
			Range(1, 5, 1),
			FromSlice([]bool{true, true, true, true}),
			[]int{1, 2, 3, 4},
		},
		{
			"only false",
			Range(1, 5, 1),
			FromSlice([]bool{false, false, false, false}),
			[]int{},
		},
		{
			"yield only if cond is true",
			Range(1, 5, 1),
			FromSlice([]bool{true, false, true, false}),
			[]int{1, 3},
		},
		{
			"gen is shorter",
			Range(1, 4, 1),
			FromSlice([]bool{true, true, true, true}),
			[]int{1, 2, 3},
		},
		{
			"cond is shorter",
			Range(1, 5, 1),
			FromSlice([]bool{true, true, true}),
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Where(tt.gen, tt.cond)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestBind(t *testing.T) {
	tests := []struct {
		name     string
		gen      Gen[int]
		f        func(int) Gen[int]
		expected []int
	}{
		{
			"id",
			Range(1, 5, 1),
			func(i int) Gen[int] { return Const(i) },
			[]int{1, 2, 3, 4},
		},
		{
			"twice",
			Range(1, 5, 1),
			func(i int) Gen[int] { return Add(Const(i), Const(i)) },
			[]int{2, 4, 6, 8},
		},
		{
			"add one",
			Range(1, 5, 1),
			func(i int) Gen[int] { return Const(i + 1) },
			[]int{2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Bind(tt.gen, tt.f)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIf(t *testing.T) {
	tests := []struct {
		name     string
		condG    Gen[bool]
		thenG    Gen[int]
		elseG    Gen[int]
		expected []int
	}{
		{
			"then",
			Const(true),
			Const(1),
			Const(2),
			[]int{1},
		},
		{
			"else",
			Const(false),
			Const(1),
			Const(2),
			[]int{2},
		},
		{
			"multiple",
			FromSlice([]bool{true, false}),
			FromSlice([]int{1, 2}),
			FromSlice([]int{3, 4}),
			[]int{1, 4},
		},
		{
			"cond is shorter",
			FromSlice([]bool{true}),
			FromSlice([]int{1, 2}),
			FromSlice([]int{3, 4}),
			[]int{1},
		},
		{
			"then is shorter",
			FromSlice([]bool{true, false}),
			FromSlice([]int{1}),
			FromSlice([]int{3, 4}),
			[]int{1},
		},
		{
			"else is shorter",
			FromSlice([]bool{true, false}),
			FromSlice([]int{1, 2}),
			FromSlice([]int{3}),
			[]int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := If(tt.condG, tt.thenG, tt.elseG)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name     string
		gen      Gen[bool]
		expected []bool
	}{
		{
			"with false",
			FromSlice([]bool{true, false}),
			[]bool{false},
		},
		{
			"without false",
			FromSlice([]bool{true, true}),
			[]bool{true},
		},
		{
			"empty",
			FromSlice([]bool{}),
			[]bool{true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := All(tt.gen)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name     string
		gen      Gen[bool]
		expected []bool
	}{
		{
			"with true",
			FromSlice([]bool{true, false}),
			[]bool{true},
		},
		{
			"without true",
			FromSlice([]bool{false, false}),
			[]bool{false},
		},
		{
			"empty",
			FromSlice([]bool{}),
			[]bool{false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Any(tt.gen)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLoop(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		gen := Loop(5)
		seq := gen()

		next, stop := iter.Pull(seq)
		defer stop()

		v, ok := next()
		assert.Equal(t, 5, v)
		assert.True(t, ok)

		v, ok = next()
		assert.Equal(t, 5, v)
		assert.True(t, ok)

		v, ok = next()
		assert.Equal(t, 5, v)
		assert.True(t, ok)
	})

	t.Run("string", func(t *testing.T) {
		gen := Loop("foo")
		seq := gen()

		next, stop := iter.Pull(seq)
		defer stop()

		v, ok := next()
		assert.Equal(t, "foo", v)
		assert.True(t, ok)

		v, ok = next()
		assert.Equal(t, "foo", v)
		assert.True(t, ok)

		v, ok = next()
		assert.Equal(t, "foo", v)
		assert.True(t, ok)
	})
}
