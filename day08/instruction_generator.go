package day08

import "fmt"

var (
	finalInstruction = Instruction{'Z', 'Z', 'Z'}
	buf              [9]byte
)

func NewInstructionGenerator() InstructionGenerator {
	return InstructionGenerator{
		ptr:  nil,
		data: make(map[Instruction]InstructionPair),
	}
}

type InstructionGenerator struct {
	ptr  *Instruction
	data map[Instruction]InstructionPair
}

func (m *InstructionGenerator) Add(line []byte) {
	// TBR = (RSM, NBX)

	offset := 0
	for _, symbol := range line {
		if !isCapitalLetter(symbol) {
			continue
		}

		buf[offset] = symbol
		offset++
	}

	key := Instruction{buf[0], buf[1], buf[2]}
	left := Instruction{buf[3], buf[4], buf[5]}
	right := Instruction{buf[6], buf[7], buf[8]}

	m.data[key] = InstructionPair{
		left:  left,
		right: right,
	}

	if m.ptr == nil {
		m.ptr = &key
	}
}

func (m *InstructionGenerator) Next(dir Direction) Instruction {
	val := m.data[*m.ptr]

	if dir == DirectionLeft {
		m.ptr = &val.left
		return val.left
	} else {
		m.ptr = &val.right
		return val.right
	}
}

func (m *InstructionGenerator) NextStr(dir Direction) string {
	return fmt.Sprintf("%s", m.Next(dir))
}

// Instruction can be represented as one 16bit digit (5 bits per number)
type Instruction = [3]byte

type InstructionPair struct {
	left  Instruction
	right Instruction
}

func isCapitalLetter(symbol byte) bool {
	return symbol >= 'A' && symbol <= 'Z'
}
