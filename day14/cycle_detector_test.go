package day14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCycleDetector(t *testing.T) {
	cd := NewCycleDetector()
	cd.AddAndTryDetectCycle(1, 2, 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			cd.AddAndTryDetectCycle(j)
		}
	}

	assert.Equal(t, 1, cd.Predict(0))
	assert.Equal(t, 2, cd.Predict(1))
	assert.Equal(t, 3, cd.Predict(2))
	assert.Equal(t, 0, cd.Predict(3))
	assert.Equal(t, 1, cd.Predict(4))
}
