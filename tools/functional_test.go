package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvery(t *testing.T) {
	assert.Equal(t, true, Every([]int{1, 1, 1, 1}, func(i int) bool {
		return i == 1
	}))

	assert.Equal(t, true, Every([]int{}, func(i int) bool {
		return i == 1
	}))
}
