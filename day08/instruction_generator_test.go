package day08

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLcm(t *testing.T) {

	assert.Equal(t, 6, lcm(2, 3))
	assert.Equal(t, 10, lcm(2, 5))
	assert.Equal(t, 15, lcm(3, 5))

}
