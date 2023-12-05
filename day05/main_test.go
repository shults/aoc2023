package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeSplitWhenEquals(t *testing.T) {
	r1 := NewRange(0, 9)
	r2 := NewRange(0, 9)

	res := r1.split(r2)
	assert.Equal(t, []Range{{0, 9}}, res)
}

func TestRangeSplitWhenOneContainsOther(t *testing.T) {
	r1 := NewRange(0, 15)
	r2 := NewRange(5, 10)

	assert.Equal(t, []Range{{0, 4}, {5, 10}, {11, 15}}, r1.split(r2))
	assert.Equal(t, []Range{{0, 4}, {5, 10}, {11, 15}}, r2.split(r1))
}

func TestRangeSplitWhenLeftOverlap(t *testing.T) {
	r1 := NewRange(0, 5)
	r2 := NewRange(5, 10)

	assert.Equal(t, []Range{{0, 4}, {5, 5}, {6, 10}}, r1.split(r2))
	assert.Equal(t, []Range{{0, 4}, {5, 5}, {6, 10}}, r2.split(r1))
}
