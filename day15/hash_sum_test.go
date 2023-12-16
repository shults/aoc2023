package day15

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashSum(t *testing.T) {

	assert.Equal(t, 52, Hash("HASH"))
}

func TestHashSum2(t *testing.T) {
	boxes := NewBoxes()
	boxes.ProcessInput(bytes.NewBufferString("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))

	assert.Equal(t, 1320, boxes.CalculateSumOfHashes())
}
