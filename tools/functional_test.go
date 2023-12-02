package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvery(t *testing.T) {
	isOne := func(i *int) bool {
		return *i == 1
	}

	assert.Equal(t, true, MustEvery([]int{1, 1, 1, 1}, isOne))

	assert.Equal(t, true, MustEvery([]int{}, isOne))
}
