package day16

import (
	"aoc2023/tools"
	"github.com/fatih/color"
	"strings"
)

type Direction byte

const (
	dirTop Direction = iota
	dirRight
	dirBottom
	dirLeft
)

type Tile interface {
	split(vector Direction) []PositionDirection
}

type TileSlash struct {
	Position
}

func (t TileSlash) split(vector Direction) []PositionDirection {
	pd := NewPositionDirection(t.Position, vector)

	switch vector {
	case dirTop:
		return []PositionDirection{pd.right()}
	case dirBottom:
		return []PositionDirection{pd.left()}
	case dirLeft:
		return []PositionDirection{pd.bottom()}
	case dirRight:
		return []PositionDirection{pd.top()}
	default:
		panic("unexpected")
	}
}

type TileBackSlash struct {
	Position
}

func (t TileBackSlash) split(vector Direction) []PositionDirection {
	pd := NewPositionDirection(t.Position, vector)

	switch vector {
	case dirTop:
		return []PositionDirection{pd.left()}
	case dirRight:
		return []PositionDirection{pd.bottom()}
	case dirBottom:
		return []PositionDirection{pd.right()}
	case dirLeft:
		return []PositionDirection{pd.top()}
	default:
		panic("unexpected")
	}
}

type TileVertical struct {
	Position
}

func (t TileVertical) split(vector Direction) []PositionDirection {
	pd := NewPositionDirection(t.Position, vector)

	switch vector {
	case dirBottom:
		return []PositionDirection{pd.bottom()}
	case dirTop:
		return []PositionDirection{pd.top()}
	case dirRight, dirLeft:
		return []PositionDirection{
			pd.top(),
			pd.bottom(),
		}
	default:
		panic("unexpected")
	}
}

type TileHorizontal struct {
	Position
}

func (t TileHorizontal) split(vector Direction) []PositionDirection {
	pd := NewPositionDirection(t.Position, vector)

	switch vector {
	case dirBottom, dirTop:
		return []PositionDirection{pd.left(), pd.right()}
	case dirRight:
		return []PositionDirection{pd.right()}
	case dirLeft:
		return []PositionDirection{pd.left()}
	default:
		panic("unexpected")
	}
}

type TileDot struct {
	Position
}

func (t TileDot) split(vector Direction) []PositionDirection {
	return []PositionDirection{
		NewPositionDirection(t.Position, vector).next(vector),
	}
}

func NewTile(b byte, pos Position) Tile {
	switch b {
	case '.':
		return &TileDot{pos}
	case '-':
		return &TileHorizontal{pos}
	case '|':
		return &TileVertical{pos}
	case '/':
		return &TileSlash{pos}
	case '\\':
		return &TileBackSlash{pos}
	default:
		panic("unexpected case")
	}
}

type BeamSplitter struct {
	tiles      [][]Tile
	rows, cols int
	values     [][]byte
}

func NewBeamSplitter(rows []string) *BeamSplitter {
	tiles := make([][]Tile, len(rows))
	values := make([][]byte, len(rows))

	for i, row := range rows {
		tiles[i] = make([]Tile, len([]byte(row)))
		values[i] = []byte(row)
		for j, symbol := range []byte(row) {
			tiles[i][j] = NewTile(symbol, Position{i, j})
		}
	}

	return &BeamSplitter{
		tiles:  tiles,
		values: values,
		rows:   len(tiles),
		cols:   len(tiles[0]),
	}
}

func (b *BeamSplitter) CalculateEnergizedTiles(verbose bool) int {
	//cache := make(map[PositionDirection]int)
	visitSet := NewVisitSet()
	startPosition := NewPositionDirection(Position{0, 0}, dirRight)

	val := b.traverse(
		startPosition,
		visitSet,
	)

	return val
}

func (b *BeamSplitter) CalculateMaxEnergizedTiles(verbose bool) int {
	//cache := make(map[PositionDirection]int)
	dim := tools.MatrixSize(b.tiles)

	maxValue := 0

	for _, pair := range []struct {
		dir Direction
		gen func(i int) Position
	}{
		{dirBottom, func(i int) Position {
			return Position{0, i}
		}},
		{dirTop, func(i int) Position {
			return Position{dim.Rows - 1, i}
		}},
		{dirRight, func(i int) Position {
			return Position{i, 0}
		}},
		{dirLeft, func(i int) Position {
			return Position{i, dim.Cols - 1}
		}},
	} {
		for i := 0; i < tools.Max(dim.Cols, dim.Rows); i++ {
			startPosition := NewPositionDirection(pair.gen(i), pair.dir)
			visitSet := NewVisitSet()

			maxValue = tools.Max(maxValue, b.traverse(
				startPosition,
				visitSet,
			))
		}
	}

	//val, ok := cache[startPosition]

	return maxValue
}

func (b *BeamSplitter) ToString(energizedTiles map[Position]struct{}) string {
	yellow := color.New(color.BgYellow)

	var sb strings.Builder
	for i := 0; i < len(b.tiles); i++ {
		for j := 0; j < len(b.tiles[i]); j++ {
			_, wasVisited := energizedTiles[Position{i, j}]

			if wasVisited {
				sb.Write([]byte(yellow.Sprintf("%s", []byte{b.values[i][j]})))
			} else {
				sb.WriteByte(b.values[i][j])
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func (b *BeamSplitter) traverse(
	pd PositionDirection,
	visitSet *VisitSet,
) int {
	if !b.isInside(pd.Position) {
		return 0
	}

	if visitSet.Has(pd) {
		return 0
	}

	visitSet.Add(pd)
	tile := b.tiles[pd.i][pd.j]
	nextDirections := tile.split(pd.Direction)

	for _, pdNext := range nextDirections {
		b.traverse(pdNext, visitSet)
	}

	return visitSet.UniqPositions()
}

func (b *BeamSplitter) isInside(pos Position) bool {
	return pos.i >= 0 && pos.j >= 0 && pos.i < b.rows && pos.j < b.cols
}

type Position struct {
	i, j int
}

type PositionDirection struct {
	Position
	Direction
}

func NewPositionDirection(pos Position, dir Direction) PositionDirection {
	return PositionDirection{
		Position:  pos,
		Direction: dir,
	}
}

func (pd PositionDirection) top() PositionDirection {
	pd.Direction = dirTop
	pd.Position.i--
	return pd
}

func (pd PositionDirection) bottom() PositionDirection {
	pd.Direction = dirBottom
	pd.Position.i++
	return pd
}

func (pd PositionDirection) left() PositionDirection {
	pd.Direction = dirLeft
	pd.Position.j--
	return pd
}

func (pd PositionDirection) right() PositionDirection {
	pd.Direction = dirRight
	pd.Position.j++
	return pd
}

func (pd PositionDirection) next(dir Direction) PositionDirection {
	switch dir {
	case dirLeft:
		return pd.left()
	case dirRight:
		return pd.right()
	case dirTop:
		return pd.top()
	case dirBottom:
		return pd.bottom()
	default:
		panic("unexpected case")
	}
}

type VisitSet struct {
	posDirMap map[PositionDirection]struct{}
	posMap    map[Position]struct{}
}

func NewVisitSet() *VisitSet {
	return &VisitSet{
		posMap:    make(map[Position]struct{}),
		posDirMap: make(map[PositionDirection]struct{}),
	}
}

func (v *VisitSet) Add(posDir PositionDirection) {
	v.posDirMap[posDir] = struct{}{}
	v.posMap[posDir.Position] = struct{}{}
}

func (v *VisitSet) Has(posDir PositionDirection) bool {
	_, has := v.posDirMap[posDir]
	return has
}

func (v *VisitSet) HasPosition(pos Position) bool {
	_, has := v.posMap[pos]
	return has
}

func (v *VisitSet) UniqPositions() int {
	return len(v.posMap)
}
