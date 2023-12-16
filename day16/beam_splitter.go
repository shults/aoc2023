package day16

import (
	"fmt"
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
	visited := make(map[PositionDirection]struct{})
	b.traverse(NewPositionDirection(Position{0, 0}, dirRight), visited)

	energizedTiles := make(map[Position]struct{})

	for v := range visited {
		energizedTiles[v.Position] = struct{}{}
	}

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

	if verbose {
		fmt.Printf("%s\b", sb.String())
	}

	return len(energizedTiles)
}

func (b *BeamSplitter) traverse(pd PositionDirection, visited map[PositionDirection]struct{}) {
	if !b.isInside(pd.Position) {
		return
	}

	if _, alreadyVisited := visited[pd]; alreadyVisited {
		return
	}

	visited[pd] = struct{}{}

	tile := b.tiles[pd.i][pd.j]
	nextDirections := tile.split(pd.Direction)

	for _, pdNext := range nextDirections {
		b.traverse(pdNext, visited)
	}

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
