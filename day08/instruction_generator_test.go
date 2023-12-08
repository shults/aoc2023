package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstructionGenerator(t *testing.T) {
	lines := []string{
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)",
		"EEE = (EEE, EEE)",
		"GGG = (GGG, GGG)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	gen := NewInstructionGenerator(lines)

	assert.Equal(t, "BBB", gen.NextStr(DirectionLeft))
	assert.Equal(t, "EEE", gen.NextStr(RightSymbol))
	assert.Equal(t, "EEE", gen.NextStr(LeftSymbol))
}

func TestLcm(t *testing.T) {

	assert.Equal(t, 6, lcm(2, 3))
	assert.Equal(t, 10, lcm(2, 5))
	assert.Equal(t, 15, lcm(3, 5))

}
