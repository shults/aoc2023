package day08

import "fmt"

var (
	buf [9]byte
)

func NewInstructionGenerator(
	lines []string,
) InstructionGenerator {

	data := make([]mapItemPos, len(lines))
	posToInputValue := make([]Instruction, len(lines))

	instructionToValue := make(map[Instruction]struct {
		pos              uint16
		val, left, right Instruction
	})

	for pos, line := range lines {
		offset := 0
		for _, symbol := range []byte(line) {
			if isCapitalLetterOrNumeric(symbol) {
				buf[offset] = symbol
				offset++
			}
		}

		val := NewInstruction(buf[0], buf[1], buf[2])
		left := NewInstruction(buf[3], buf[4], buf[5])
		right := NewInstruction(buf[6], buf[7], buf[8])

		instructionToValue[val] = tmpStruct{
			pos:   uint16(pos),
			val:   val,
			left:  left,
			right: right,
		}
	}

	for _, mapItem := range instructionToValue {
		left := instructionToValue[mapItem.left]
		right := instructionToValue[mapItem.right]

		data[mapItem.pos] = mapItemPos{
			left:        left.pos,
			right:       right.pos,
			Instruction: mapItem.val,
		}

		posToInputValue[mapItem.pos] = mapItem.val
	}

	finalInstruction := instructionToValue[NewInstruction('Z', 'Z', 'Z')]
	firstInstruction := instructionToValue[NewInstruction('A', 'A', 'A')]

	return InstructionGenerator{
		pos:              firstInstruction.pos,
		data:             data,
		inMap:            instructionToValue,
		finalInstruction: finalInstruction.pos,
		posToInputValue:  posToInputValue,
	}
}

type InstructionGenerator struct {
	pos              uint16
	data             []mapItemPos
	posToInputValue  []Instruction
	finalInstruction uint16
	inMap            map[Instruction]struct {
		pos              uint16
		val, left, right Instruction
	}
}

type tmpStruct struct {
	pos              uint16
	val, left, right Instruction
}

type mapItemPos struct {
	left, right uint16
	Instruction
}

func (p *mapItemPos) GetNext(dir Direction) uint16 {
	if dir == DirectionLeft {
		return p.left
	} else {
		return p.right
	}
}

func (m *InstructionGenerator) Next(dir Direction) MappedInstruction {
	val := m.data[m.pos]

	if dir == DirectionLeft {
		m.pos = val.left
		return val.left
	} else {
		m.pos = val.right
		return val.right
	}
}

func (m *InstructionGenerator) NextBulk(dir Direction, positions []int) MappedInstruction {
	val := m.data[m.pos]

	if dir == DirectionLeft {
		m.pos = val.left
		return val.left
	} else {
		m.pos = val.right
		return val.right
	}
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	for {
		if b == 0 {
			break
		}

		var r int
		if a > b {
			r = a % b
		} else {
			r = b % a
		}
		a = b
		b = r
	}

	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func (m *InstructionGenerator) GenFinialState(pos uint16, dirGen DirectionGenerator) (result int) {
	result = 0

	for {
		node := m.data[pos]
		pos = node.GetNext(dirGen.Next())

		if node.isFinal {
			return
		}

		result++
	}
}

func (m *InstructionGenerator) Part2(dirGen DirectionGenerator, verbose bool) (iters int) {
	dirGen.Reset()
	var positions []uint16

	for pos := 0; pos < len(m.data); pos++ {
		node := m.data[pos]
		if node.isStart {
			positions = append(positions, uint16(pos))
		}
	}

	iters = 0

	results := make([]int, len(positions))

	for i, pos := range positions {
		results[i] = m.GenFinialState(pos, dirGen.CloneAndReset())
	}

	if verbose {
		fmt.Printf("%+v\n", results)
	}

	iters = results[0]
	for i := 1; i < len(results); i++ {
		iters = lcm(iters, results[i])
	}

	return
}

func (m *InstructionGenerator) NextStr(dir Direction) string {
	// todo: make slice which has mapped pos to value
	return fmt.Sprintf("%s", m.posToInputValue[m.Next(dir)].val)
}

type MappedInstruction = uint16

// Instruction can be represented as one 16bit digit (5 bits per number)

func NewInstruction(a, b, c byte) Instruction {
	return Instruction{
		val:     [3]byte{a, b, c},
		isFinal: c == 'Z',
		isStart: c == 'A',
	}
}

type Instruction struct {
	val     [3]byte
	isFinal bool
	isStart bool
}

type InstructionPair struct {
	left  Instruction
	right Instruction
}

func isCapitalLetterOrNumeric(symbol byte) bool {
	return (symbol >= 'A' && symbol <= 'Z') || (symbol >= '0' && symbol <= '9')
}
