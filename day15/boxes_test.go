package day15

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInstruction(t *testing.T) {
	ins1 := NewInstruction("abc-")
	assert.Equal(t, "abc-", ins1.String())

	ins2 := NewInstruction("abc=23")
	assert.Equal(t, "abc=23", ins2.String())
}

func TestBoxes_CalculateFocalSum(t *testing.T) {
	boxes := NewBoxes()
	boxes.ProcessInput(bytes.NewBufferString("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"))
	assert.Equal(t, 145, boxes.CalculateFocalSum())
}

func TestBoxes_Dummy(t *testing.T) {
	s := []int{1, 2, 3}

	fmt.Printf("%+v\n", s[:2])
	fmt.Printf("%+v\n", s[2:])
}
