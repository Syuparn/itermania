package itermania

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNot(t *testing.T) {
	tests := []struct {
		name     string
		x        Gen[bool]
		expected []bool
	}{
		{
			"![true]",
			Const(true),
			[]bool{false},
		},
		{
			"![false]",
			Const(false),
			[]bool{true},
		},
		{
			"![true, false]",
			FromSlice([]bool{true, false}),
			[]bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := Not(tt.x)
			actual := ToSlice(gen)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
