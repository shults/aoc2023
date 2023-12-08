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

const bufSize = 1 << 12

func NewDirectionGenerator(line []byte) DirectionGenerator {
	dirGen := DirectionGenerator{
		pos:  0,
		size: len(line),
	}

	if len(line) > bufSize {
		panic("too small buffer")
	}

	for key, instruction := range line {
		if instruction == LeftSymbol {
			dirGen.data[key] = DirectionLeft
		} else if instruction == RightSymbol {
			dirGen.data[key] = DirectionRight
		} else {
			panic("wrong symbol")
		}
	}

	return dirGen
}

type DirectionGenerator struct {
	pos  int
	size int
	data [bufSize]Direction
}

func (g *DirectionGenerator) Next() Direction {
	ret := g.data[g.pos]
	g.pos++

	if g.pos == g.size {
		g.pos = 0
	}

	return ret
}

func (g *DirectionGenerator) Reset() {
	g.pos = 0
}

func (g *DirectionGenerator) CloneAndReset() DirectionGenerator {
	clone := *g
	clone.Reset()
	return clone
}
