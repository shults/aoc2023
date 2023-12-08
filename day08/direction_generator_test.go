package day08

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInstructionGenerator(t *testing.T) {
	gen := NewDirectionGenerator([]byte("LLRRL"))

	const iterations = 10
	tests := []Direction{
		DirectionLeft,
		DirectionLeft,
		DirectionRight,
		DirectionRight,
		DirectionLeft,
	}

	for i := 0; i < iterations; i++ {
		for _, test := range tests {
			val := gen.Next()
			assert.Equal(t, test, val, fmt.Sprintf("expected=%d got=%d", test, val))
		}
	}

	assert.Equal(t, uint64(0), gen.pos)
	assert.Equal(t, byte(DirectionLeft), gen.Next())
	assert.Equal(t, uint64(1), gen.pos)
	assert.Equal(t, byte(DirectionLeft), gen.Next())
	assert.Equal(t, byte(DirectionRight), gen.Next())
	assert.Equal(t, byte(DirectionRight), gen.Next())
	assert.Equal(t, uint64(4), gen.pos)
	assert.Equal(t, byte(DirectionLeft), gen.Next())
	assert.Equal(t, uint64(0), gen.pos)
}
