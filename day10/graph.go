package day10

import "fmt"

func NewAdjacentHelper(matrix [][]*Tile) *AdjacentHelper {

	return &AdjacentHelper{
		maxRow: len(matrix) - 1,
		maxCol: len(matrix[0]) - 1,
		matrix: matrix,
	}
}

type AdjacentHelper struct {
	maxRow, maxCol int
	matrix         [][]*Tile
}

func (a *AdjacentHelper) CanAddToAdjacent(currentTile *Tile, moveVector Direction) (bool, *Tile) {
	nextPosition := currentTile.nextPosition(moveVector)

	if !a.contains(nextPosition) {
		return false, nil
	}

	nextTile := a.matrix[nextPosition.row][nextPosition.col]

	if currentTile.isStartTile {
		return nextTile.CanBeVisitedFrom(moveVector.opposite()), nextTile
	}

	// if I cannot be visited from opposite direction
	if currentTile.CanBeVisitedFrom(moveVector) && nextTile.CanBeVisitedFrom(moveVector.opposite()) {
		// i cant be visited nextPosition opposite direction
		return true, nextTile
	}

	return false, nil
}

func (a *AdjacentHelper) contains(p Position) bool {
	return p.row >= 0 && p.col >= 0 && p.row <= a.maxRow && p.col <= a.maxCol
}

type Position struct {
	row, col int
}

func NewPosition(row, col int) Position {
	return Position{
		row: row,
		col: col,
	}
}

func (p Position) nextPosition(dir Direction) Position {
	if dir.isNorth() {
		return p.north()
	} else if dir.isEast() {
		return p.east()
	} else if dir.isSouth() {
		return p.south()
	} else if dir.isWest() {
		return p.west()
	} else {
		panic("unknown direction")
	}
}

func (p Position) west() Position {
	p.col--
	return p
}

func (p Position) east() Position {
	p.col++
	return p
}

func (p Position) north() Position {
	p.row--
	return p
}

func (p Position) south() Position {
	p.row++
	return p
}

type Tile struct {
	Position
	TileSymbol
	adjacentNodes []*Tile
	isStartTile   bool
}

func NewTile(pos Position, sym TileSymbol, tileSymbol byte) Tile {
	return Tile{
		Position:      pos,
		TileSymbol:    sym,
		adjacentNodes: nil,
		isStartTile:   tileSymbol == start,
	}
}

func (t *Tile) BuildAdjacent(adjacentHelper *AdjacentHelper) {
	// if I as not can be visited nextPosition North direction and my North neighbour can be visited nextPosition South than I can add him to adjacent nodes
	for _, dir := range []Direction{
		directionNorth,
		directionEast,
		directionSouth,
		directionWest,
	} {
		if ok, node := adjacentHelper.CanAddToAdjacent(t, dir); ok {
			t.adjacentNodes = append(t.adjacentNodes, node)
		}
	}
}

func (t *Tile) TraverseInDepth(visitedSet *VisitedSet) (nodesToVisit []*Tile) {
	if visitedSet.WasVisited(t.Position) {
		return []*Tile{}
	}

	visitedSet.Add(t.Position)

	nodesToVisit = make([]*Tile, 0)

	for _, adjacentNode := range t.adjacentNodes {
		if !visitedSet.WasVisited(adjacentNode.Position) {
			nodesToVisit = append(nodesToVisit, adjacentNode)
		}
	}

	return nodesToVisit
}

type TileSymbol interface {
	CanBeVisitedFrom(from Direction) bool
}

var _ TileSymbol = (*StartTile)(nil)

type StartTile struct{}

func (t StartTile) CanBeVisitedFrom(from Direction) bool {
	return false
}

var _ TileSymbol = (*GroundTile)(nil)

type GroundTile struct{}

func (g GroundTile) CanBeVisitedFrom(from Direction) bool {
	return false
}

var _ TileSymbol = (*SouthWestTile)(nil)

type SouthWestTile struct{}

func (s SouthWestTile) CanBeVisitedFrom(from Direction) bool {
	return from.isSouth() || from.isWest()
}

var _ TileSymbol = (*SouthEastTile)(nil)

type SouthEastTile struct{}

func (s SouthEastTile) CanBeVisitedFrom(from Direction) bool {
	return from.isSouth() || from.isEast()
}

var _ TileSymbol = (*NorthEastTile)(nil)

type NorthEastTile struct{}

func (s NorthEastTile) CanBeVisitedFrom(from Direction) bool {
	return from.isNorth() || from.isEast()
}

var _ TileSymbol = (*NorthWestTile)(nil)

type NorthWestTile struct{}

func (s NorthWestTile) CanBeVisitedFrom(from Direction) bool {
	return from.isNorth() || from.isWest()
}

var _ TileSymbol = (*HorizontalTile)(nil)

type HorizontalTile struct{}

func (s HorizontalTile) CanBeVisitedFrom(from Direction) bool {
	return from.isEast() || from.isWest()
}

var _ TileSymbol = (*VerticalTile)(nil)

type VerticalTile struct{}

func (s VerticalTile) CanBeVisitedFrom(from Direction) bool {
	return from.isNorth() || from.isSouth()
}

func NewTileSymbol(val byte) TileSymbol {
	switch val {
	case vertical:
		return &VerticalTile{}
	case horizontal:
		return &HorizontalTile{}
	case northEast:
		return &NorthEastTile{}
	case northWest:
		return &NorthWestTile{}
	case southEast:
		return &SouthEastTile{}
	case southWest:
		return &SouthWestTile{}
	case ground:
		return &GroundTile{}
	case start:
		return &StartTile{}
	}

	panic(fmt.Errorf("unexpected symbol '%s'", []byte{val}))
}

// can go to?
// visited map

const (
	// is a vertical pipe connecting north and south.
	vertical byte = '|'
	// is a horizontal pipe connecting east and west.
	horizontal byte = '-'
	// L is a 90-degree bend connecting north and east.
	northEast byte = 'L'
	// J is a 90-degree bend connecting north and west.
	northWest byte = 'J'
	// 7 is a 90-degree bend connecting south and west.
	southWest byte = '7'
	// F is a 90-degree bend connecting south and east.
	southEast byte = 'F'
	// . is ground; there is no pipe in this tile.
	ground byte = '.'
	// S is the starting position of the animal; there is a pipe on this
	start byte = 'S'
)

type Direction struct {
	symbol DirectionSymbol
}

func NewDirection(symbol DirectionSymbol) Direction {
	switch symbol {
	case dirWest:
	case dirEast:
	case dirNorth:
	case dirSouth:

	default:
		panic(fmt.Errorf("unknown direction '%s'", []byte{byte(symbol)}))
	}

	return Direction{
		symbol: symbol,
	}
}

func (d Direction) opposite() Direction {
	switch d.symbol {
	case dirWest:
		return directionEast
	case dirEast:
		return directionWest
	case dirNorth:
		return directionSouth
	case dirSouth:
		return directionNorth
	default:
		panic(fmt.Errorf("unknown direction '%s'", []byte{byte(d.symbol)}))
	}
}

func (d Direction) isWest() bool {
	return d.symbol == dirWest
}

func (d Direction) isEast() bool {
	return d.symbol == dirEast
}

func (d Direction) isNorth() bool {
	return d.symbol == dirNorth
}

func (d Direction) isSouth() bool {
	return d.symbol == dirSouth
}

type DirectionSymbol byte

const (
	dirNorth DirectionSymbol = 'N'
	dirSouth DirectionSymbol = 'S'
	dirWest  DirectionSymbol = 'W'
	dirEast  DirectionSymbol = 'E'
)

var (
	directionNorth = NewDirection(dirNorth)
	directionSouth = NewDirection(dirSouth)
	directionWest  = NewDirection(dirWest)
	directionEast  = NewDirection(dirEast)
)

type Graph struct {
	startTile *Tile
}

func NewGraph(field []string) Graph {
	tilesMatrix := make([][]*Tile, len(field))
	var startTile *Tile

	for row, rawSymbols := range field {
		positions := make([]*Tile, len(rawSymbols))

		for col, tileSymbol := range []byte(rawSymbols) {
			tile := NewTile(
				NewPosition(row, col),
				NewTileSymbol(tileSymbol),
				tileSymbol,
			)

			if tile.isStartTile {
				startTile = &tile
			}

			positions[col] = &tile
		}

		tilesMatrix[row] = positions
	}

	adjacentHelper := NewAdjacentHelper(tilesMatrix)

	for _, tileRow := range tilesMatrix {
		for _, tile := range tileRow {
			tile.BuildAdjacent(adjacentHelper)
		}
	}

	return Graph{
		startTile: startTile,
	}
}

func (g *Graph) CalculatePart1() int {
	set := NewVisitedSet()
	nodesToVisit := []*Tile{g.startTile}

	iters := 0
	for {
		newNodesToVisit := make([]*Tile, 0)

		for _, node := range nodesToVisit {
			newNodesToVisit = append(newNodesToVisit, node.TraverseInDepth(set)...)
		}

		if len(newNodesToVisit) == 0 {
			break
		}

		nodesToVisit = newNodesToVisit
		iters++
	}

	return iters
}

type VisitedSet struct {
	inner map[Position]struct{}
}

func NewVisitedSet() *VisitedSet {
	return &VisitedSet{
		inner: make(map[Position]struct{}),
	}
}

func (s *VisitedSet) WasVisited(p Position) bool {
	_, wasVisited := s.inner[p]
	return wasVisited
}

func (s *VisitedSet) Add(p Position) {
	s.inner[p] = struct{}{}
}
