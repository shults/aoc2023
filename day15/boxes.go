package day15

import (
	"aoc2023/tools"
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Boxes struct {
	boxes        []*Box
	instructions []Instruction
}

func NewBoxes() *Boxes {
	boxes := make([]*Box, 256)

	for i, _ := range boxes {
		boxes[i] = NewBox(i + 1)
	}

	return &Boxes{
		boxes: boxes,
	}
}

func (b *Boxes) ProcessInput(in io.Reader) {
	reader := bufio.NewReader(in)

	for {
		ins, err := reader.ReadString(',')

		if err == io.EOF {
			b.processInstruction(NewInstruction(ins))
			break
		}

		if err != nil {
			panic(err)
		}

		b.processInstruction(NewInstruction(ins[:len(ins)-1]))
	}
}

func (b *Boxes) processInstruction(ins Instruction) {
	b.instructions = append(b.instructions, ins)

	if ins.IsInsert() {
		b.boxes[ins.labelHash].Insert(ins)
	} else {
		b.boxes[ins.labelHash].Remove(ins)
	}
}

func (b *Boxes) CalculateFocalSum() int {
	focalSum := 0

	for _, box := range b.boxes {
		focalSum += box.calculateFocalSum()
	}

	return focalSum
}

func (b *Boxes) CalculateSumOfHashes() int {
	sumOfHashes := 0

	for _, instruction := range b.instructions {
		sumOfHashes += instruction.instructionHash
	}

	return sumOfHashes
}

type Box struct {
	nr    int
	slots []Lens
}

func NewBox(nr int) *Box {
	return &Box{
		nr:    nr,
		slots: make([]Lens, 0),
	}
}

func (b *Box) Insert(ins Instruction) {
	newLens := Lens{
		label:       ins.label,
		focalLength: ins.val,
	}

	for i, slot := range b.slots {
		if slot.label == newLens.label {
			b.slots[i] = newLens
			return
		}
	}

	b.slots = append(b.slots, newLens)
}

func (b *Box) Remove(ins Instruction) {
	for i, slot := range b.slots {
		if slot.label == ins.label {
			b.slots = append(b.slots[:i], b.slots[i+1:]...)

			// If the operation character is a dash (-), go to the relevant box and remove the lens with the given label if it is present in the box.
			// Then, move any remaining lenses as far forward in the box as they can go without changing their order, filling any space made by removing the indicated lens.
			// (If no lens in that box has the given label, nothing happens.)

			return
		}
	}
}

func (b *Box) calculateFocalSum() int {
	focalSum := 0

	for i, slot := range b.slots {
		focalSum += b.nr * (i + 1) * slot.focalLength
	}

	return focalSum
}

type LensLabel string

type Lens struct {
	label       LensLabel
	focalLength int
}

type Instruction struct {
	label           LensLabel
	operation       byte
	val             int
	instructionHash int
	labelHash       int
}

func NewInstruction(instruction string) Instruction {
	var labelLen, val int

	for i, symbol := range []byte(instruction) {
		labelLen = i
		if symbol == '-' || symbol == '=' {
			break
		}
	}

	if instruction[labelLen] == '=' {
		var err error
		val, err = strconv.Atoi(instruction[labelLen+1:])
		tools.AssertNoError(err)
	}

	return Instruction{
		label:           LensLabel(instruction[:labelLen]),
		operation:       instruction[labelLen],
		val:             val,
		instructionHash: Hash(instruction),
		labelHash:       Hash(instruction[:labelLen]),
	}
}

func (i *Instruction) IsInsert() bool {
	return i.operation == '='
}

func (i *Instruction) IsRemove() bool {
	return i.operation == '-'
}

func (i *Instruction) String() string {
	if i.IsInsert() {
		return fmt.Sprintf("%s%s%d", i.label, []byte{i.operation}, i.val)
	} else {
		return fmt.Sprintf("%s%s", i.label, []byte{i.operation})
	}
}
