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
