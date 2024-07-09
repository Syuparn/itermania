package itermania

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[bool]
		y        Gen[bool]
		expected []bool
	}{
		{
			"[true] && [true]",
			Const(true),
			Const(true),
			[]bool{true},
		},
		{
			"[true] && [false]",
			Const(true),
			Const(false),
			[]bool{false},
		},
		{
			"[false] && [true]",
			Const(false),
			Const(true),
			[]bool{false},
		},
		{
			"[false] && [false]",
			Const(false),
			Const(false),
			[]bool{false},
		},
		{
			"[true, false] && [true]",
			FromSlice([]bool{true, false}),
			Const(true),
			[]bool{true, false},
		},
		{
			"[true] && [false, true]",
			Const(true),
			FromSlice([]bool{false, true}),
			[]bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := And(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[bool]
		y        Gen[bool]
		expected []bool
	}{
		{
			"[true] || [true]",
			Const(true),
			Const(true),
			[]bool{true},
		},
		{
			"[true] || [false]",
			Const(true),
			Const(false),
			[]bool{true},
		},
		{
			"[false] || [true]",
			Const(false),
			Const(true),
			[]bool{true},
		},
		{
			"[false] || [false]",
			Const(false),
			Const(false),
			[]bool{false},
		},
		{
			"[true, false] || [false]",
			FromSlice([]bool{true, false}),
			Const(false),
			[]bool{true, false},
		},
		{
			"[true] || [false, true]",
			Const(false),
			FromSlice([]bool{false, true}),
			[]bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Or(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestEq(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []bool
	}{
		{
			"[1] == [1]",
			Const(1),
			Const(1),
			[]bool{true},
		},
		{
			"[1] == [2]",
			Const(1),
			Const(2),
			[]bool{false},
		},
		{
			"[1, 2] == [1]",
			FromSlice([]int{1, 2}),
			Const(1),
			[]bool{true, false},
		},
		{
			"[1] || [1, 2]",
			Const(1),
			FromSlice([]int{1, 2}),
			[]bool{true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Eq(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNeq(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []bool
	}{
		{
			"[1] != [1]",
			Const(1),
			Const(1),
			[]bool{false},
		},
		{
			"[1] != [2]",
			Const(1),
			Const(2),
			[]bool{true},
		},
		{
			"[1, 2] != [1]",
			FromSlice([]int{1, 2}),
			Const(1),
			[]bool{false, true},
		},
		{
			"[1] || [1, 2]",
			Const(1),
			FromSlice([]int{1, 2}),
			[]bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Neq(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGt(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []bool
	}{
		{
			"[2] > [1]",
			Const(2),
			Const(1),
			[]bool{true},
		},
		{
			"[2] > [2]",
			Const(2),
			Const(2),
			[]bool{false},
		},
		{
			"[2] > [3]",
			Const(2),
			Const(3),
			[]bool{false},
		},
		{
			"[1, 2] > [1]",
			FromSlice([]int{1, 2}),
			Const(1),
			[]bool{false, true},
		},
		{
			"[2] > [1, 2]",
			Const(2),
			FromSlice([]int{1, 2}),
			[]bool{true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Gt(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLt(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []bool
	}{
		{
			"[2] < [1]",
			Const(2),
			Const(1),
			[]bool{false},
		},
		{
			"[2] < [2]",
			Const(2),
			Const(2),
			[]bool{false},
		},
		{
			"[2] < [3]",
			Const(2),
			Const(3),
			[]bool{true},
		},
		{
			"[1, 2] < [2]",
			FromSlice([]int{1, 2}),
			Const(2),
			[]bool{true, false},
		},
		{
			"[1] < [1, 2]",
			Const(1),
			FromSlice([]int{1, 2}),
			[]bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Lt(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGe(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []bool
	}{
		{
			"[2] >= [1]",
			Const(2),
			Const(1),
			[]bool{true},
		},
		{
			"[2] >= [2]",
			Const(2),
			Const(2),
			[]bool{true},
		},
		{
			"[2] >= [3]",
			Const(2),
			Const(3),
			[]bool{false},
		},
		{
			"[1, 2] >= [2]",
			FromSlice([]int{1, 2}),
			Const(2),
			[]bool{false, true},
		},
		{
			"[1] >= [1, 2]",
			Const(1),
			FromSlice([]int{1, 2}),
			[]bool{true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Ge(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLe(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []bool
	}{
		{
			"[2] <= [1]",
			Const(2),
			Const(1),
			[]bool{false},
		},
		{
			"[2] <= [2]",
			Const(2),
			Const(2),
			[]bool{true},
		},
		{
			"[2] <= [3]",
			Const(2),
			Const(3),
			[]bool{true},
		},
		{
			"[1, 2] <= [1]",
			FromSlice([]int{1, 2}),
			Const(1),
			[]bool{true, false},
		},
		{
			"[2] <= [1, 2]",
			Const(2),
			FromSlice([]int{1, 2}),
			[]bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Le(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []int
	}{
		{
			"[2] + [3]",
			Const(2),
			Const(3),
			[]int{5},
		},
		{
			"[1, 2] + [3]",
			FromSlice([]int{1, 2}),
			Const(3),
			[]int{4, 5},
		},
		{
			"[1] + [2, 3]",
			Const(1),
			FromSlice([]int{2, 3}),
			[]int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Add(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []int
	}{
		{
			"[2] - [3]",
			Const(2),
			Const(3),
			[]int{-1},
		},
		{
			"[1, 2] - [3]",
			FromSlice([]int{1, 2}),
			Const(3),
			[]int{-2, -1},
		},
		{
			"[1] - [2, 3]",
			Const(1),
			FromSlice([]int{2, 3}),
			[]int{-1, -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Sub(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []int
	}{
		{
			"[2] * [3]",
			Const(2),
			Const(3),
			[]int{6},
		},
		{
			"[1, 2] * [3]",
			FromSlice([]int{1, 2}),
			Const(3),
			[]int{3, 6},
		},
		{
			"[1] * [2, 3]",
			Const(1),
			FromSlice([]int{2, 3}),
			[]int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Mul(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []int
	}{
		{
			"[6] / [2]",
			Const(6),
			Const(2),
			[]int{3},
		},
		{
			"[6, 4] / [2]",
			FromSlice([]int{6, 4}),
			Const(2),
			[]int{3, 2},
		},
		{
			"[6] / [2, 3]",
			Const(6),
			FromSlice([]int{2, 3}),
			[]int{3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Div(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMod(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[int]
		y        Gen[int]
		expected []int
	}{
		{
			"[6] % [5]",
			Const(6),
			Const(5),
			[]int{1},
		},
		{
			"[6, 4] % [3]",
			FromSlice([]int{6, 4}),
			Const(3),
			[]int{0, 1},
		},
		{
			"[6] % [2, 4]",
			Const(6),
			FromSlice([]int{2, 4}),
			[]int{0, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Mod(tt.x, tt.y)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
