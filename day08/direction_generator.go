package day08

type Direction = byte

const (
	DirectionLeft  Direction = 0
	DirectionRight           = 1
)

const (
	LeftSymbol  byte = 'L'
	RightSymbol      = 'R'
)

func NewDirectionGenerator(line []byte) DirectionGenerator {
	data := make([]uint64, len(line)%64+1)

	for i, instruction := range line {
		key := i / 64
		bitOffset := i % 64

		if instruction == LeftSymbol {
			continue
		}

		data[key] |= 1 << bitOffset
	}

	return DirectionGenerator{
		pos:  0,
		size: uint64(len(line)),
		data: data,
	}
}

type DirectionGenerator struct {
	pos  uint64
	size uint64
	data []uint64
}

func (g *DirectionGenerator) Next() Direction {
	key := g.pos / 64
	offset := g.pos % 64

	val := byte((g.data[key] & (1 << offset)) >> offset)

	g.pos++

	if g.pos == g.size {
		g.pos = 0
	}

	return val
}
