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

	gen := NewInstructionGenerator()

	for _, line := range lines {
		gen.Add([]byte(line))
	}

	assert.Equal(t, "BBB", gen.NextStr(DirectionLeft))
	assert.Equal(t, "EEE", gen.NextStr(RightSymbol))
	assert.Equal(t, "EEE", gen.NextStr(LeftSymbol))
}
