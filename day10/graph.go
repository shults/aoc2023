package day10

import (
	"aoc2023/tools"
	"fmt"
	"math"
)

type BoundaryHelper struct {
	minRow, maxRow, minCol, maxCol int
}

func NewBoundaryHelper(matrix [][]*Tile) BoundaryHelper {
	return BoundaryHelper{
		minCol: 0,
		minRow: 0,
		maxRow: len(matrix) - 1,
		maxCol: len(matrix[0]) - 1,
	}
}

func (b *BoundaryHelper) Contains(p Position) bool {
	return p.row >= b.minRow && p.col >= b.minCol && p.row <= b.maxRow && p.col <= b.maxCol
}

type AdjacentHelper struct {
	matrix         [][]*Tile
	BoundaryHelper BoundaryHelper
}

func NewAdjacentHelper(matrix [][]*Tile) *AdjacentHelper {
	return &AdjacentHelper{
		matrix:         matrix,
		BoundaryHelper: NewBoundaryHelper(matrix),
	}
}

func (f *AdjacentHelper) CanAddToAdjacent(currentTile *Tile, moveVector Direction) (bool, *Tile) {
	nextPosition := currentTile.nextPosition(moveVector)

	if !f.BoundaryHelper.Contains(nextPosition) {
		return false, nil
	}

	nextTile := f.matrix[nextPosition.row][nextPosition.col]

	// if I cannot be visited from opposite direction
	if currentTile.canBeVisitedFrom(moveVector, nextTile.tileSymbol) && nextTile.canBeVisitedFrom(moveVector.opposite(), currentTile.tileSymbol) {
		// i cant be visited nextPosition opposite direction
		return true, nextTile
	}

	return false, nil
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

func (p Position) direction(destination Position) Direction {
	for _, pair := range []struct {
		Position
		Direction
	}{
		{p.north(), directionNorth},
		{p.east(), directionEast},
		{p.south(), directionSouth},
		{p.west(), directionWest},
	} {
		if destination == pair.Position {
			return pair.Direction
		}
	}

	panic("not a neighbor")
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
	tileSymbol    TileSymbol
	adjacentNodes []*Tile
}

func (t *Tile) canBeVisitedFrom(from Direction, other TileSymbol) bool {
	return t.tileSymbol.canBeVisitedFrom(from, other)
}

func NewTile(pos Position, sym TileSymbol) Tile {
	return Tile{
		Position:      pos,
		tileSymbol:    sym,
		adjacentNodes: nil,
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

func (t *Tile) TraverseInWidth(visitedSet *PositionSet) (nodesToVisit []*Tile) {
	if visitedSet.Has(t.Position) {
		return []*Tile{}
	}

	visitedSet.Add(t.Position)

	nodesToVisit = make([]*Tile, 0)

	for _, adjacentNode := range t.adjacentNodes {
		if !visitedSet.Has(adjacentNode.Position) {
			nodesToVisit = append(nodesToVisit, adjacentNode)
		}
	}

	return nodesToVisit
}

func (t *Tile) TraverseInDepth(visitedSet *PositionSet, fn func(current *Tile)) {
	if visitedSet.Has(t.Position) {
		return
	}

	visitedSet.Add(t.Position)
	fn(t)

	for _, adjacentNode := range t.adjacentNodes {
		adjacentNode.TraverseInDepth(visitedSet, fn)
	}
}

type TileSymbol interface {
	canBeVisitedFrom(from Direction, other TileSymbol) bool
}

var _ TileSymbol = (*StartTile)(nil)

type StartTile struct{}

func (t StartTile) canBeVisitedFrom(from Direction, other TileSymbol) bool {
	return true
}

var _ TileSymbol = (*GroundTile)(nil)

type GroundTile struct{}

func (g GroundTile) canBeVisitedFrom(from Direction, other TileSymbol) bool {
	_, ok := other.(*GroundTile)
	return ok
}

var _ TileSymbol = (*SouthWestTile)(nil)

type SouthWestTile struct{}

func (s SouthWestTile) canBeVisitedFrom(from Direction, other TileSymbol) bool {
	return from.isSouth() || from.isWest()
}

var _ TileSymbol = (*SouthEastTile)(nil)

type SouthEastTile struct{}

func (s SouthEastTile) canBeVisitedFrom(from Direction, other TileSymbol) bool {
	return from.isSouth() || from.isEast()
}

var _ TileSymbol = (*NorthEastTile)(nil)

type NorthEastTile struct{}

func (s NorthEastTile) canBeVisitedFrom(from Direction, other TileSymbol) bool {
	return from.isNorth() || from.isEast()
}

var _ TileSymbol = (*NorthWestTile)(nil)

type NorthWestTile struct{}

func (s NorthWestTile) canBeVisitedFrom(from Direction, other TileSymbol) bool {
	return from.isNorth() || from.isWest()
}

var _ TileSymbol = (*HorizontalTile)(nil)

type HorizontalTile struct{}

func (s HorizontalTile) canBeVisitedFrom(from Direction, other TileSymbol) bool {
	return from.isEast() || from.isWest()
}

var _ TileSymbol = (*VerticalTile)(nil)

type VerticalTile struct{}

func (s VerticalTile) canBeVisitedFrom(from Direction, other TileSymbol) bool {
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
	switch d {
	case directionWest:
		return directionEast
	case directionEast:
		return directionWest
	case directionNorth:
		return directionSouth
	case directionSouth:
		return directionNorth
	default:
		panic(fmt.Errorf("unknown direction '%s'", []byte{byte(d.symbol)}))
	}
}

func (d Direction) left() Direction {
	switch d {
	case directionNorth:
		return directionWest
	case directionWest:
		return directionSouth
	case directionSouth:
		return directionEast
	case directionEast:
		return directionNorth
	default:
		panic(fmt.Errorf("unknown direction '%s'", []byte{byte(d.symbol)}))
	}
}

func (d Direction) right() Direction {
	return d.left().opposite()
}

func (d Direction) isWest() bool {
	return d == directionWest
}

func (d Direction) isEast() bool {
	return d == directionEast
}

func (d Direction) isNorth() bool {
	return d == directionNorth
}

func (d Direction) isSouth() bool {
	return d == directionSouth
}

func (d Direction) String() string {
	return fmt.Sprintf("Direction(%s)", []byte{byte(d.symbol)})
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
	startTile      *Tile
	matrix         [][]*Tile
	boundaryHelper BoundaryHelper
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
			)

			if tileSymbol == start {
				startTile = &tile
			}

			positions[col] = &tile
		}

		tilesMatrix[row] = positions
	}

	fenceHelper := NewAdjacentHelper(tilesMatrix)

	for _, tileRow := range tilesMatrix {
		for _, tile := range tileRow {
			tile.BuildAdjacent(fenceHelper)
		}
	}

	return Graph{
		startTile:      startTile,
		matrix:         tilesMatrix,
		boundaryHelper: fenceHelper.BoundaryHelper,
	}
}

func (g *Graph) CalculatePart1() int {
	set := NewPositionSet()
	nodesToVisit := []*Tile{g.startTile}

	iters := 0
	for {
		newNodesToVisit := make([]*Tile, 0)

		for _, node := range nodesToVisit {
			newNodesToVisit = append(newNodesToVisit, node.TraverseInWidth(set)...)
		}

		if len(newNodesToVisit) == 0 {
			break
		}

		nodesToVisit = newNodesToVisit
		iters++
	}

	return iters
}

func (g *Graph) CalculatePart2() int {
	fenceSet := NewPositionSet()
	fenceTiles := make([]*Tile, 0)

	g.startTile.TraverseInDepth(fenceSet, func(current *Tile) {
		if len(current.adjacentNodes) != 2 {
			panic("expected each tile to have two adjacent nodes")
		}

		fenceTiles = append(fenceTiles, current)
	})

	tools.AssertTrue(len(fenceTiles) > 0, "expected number of visited fenceTiles be greater zero")
	fenceTiles = append(fenceTiles, fenceTiles[0])

	leftSet := NewPositionSet()
	rightSet := NewPositionSet()

	for i := 0; i < len(fenceTiles)-1; i++ {
		from := fenceTiles[i]
		to := fenceTiles[i+1]
		direction := from.direction(to.Position)

		for _, tile := range []*Tile{from, to} {
			right := tile.nextPosition(direction.right())

			// is not a fence
			if !fenceSet.Has(right) {
				g.traverseInWidth(rightSet, right)
			}
		}
	}

	for i := 0; i < len(fenceTiles)-1; i++ {
		from := fenceTiles[i]
		to := fenceTiles[i+1]
		direction := from.direction(to.Position)

		for _, tile := range []*Tile{from, to} {
			left := tile.nextPosition(direction.left())
			// is not a fence
			if !fenceSet.Has(left) {
				g.traverseInWidth(leftSet, left)
			}
		}
	}

	leftSetIsOpen := g.isOpenSet(leftSet)
	rightSetIsOpen := g.isOpenSet(rightSet)

	intersection := fenceSet.Intersection(leftSet, rightSet)
	fmt.Printf("intersection %v\n", intersection)

	merged := fenceSet.Merge(leftSet, rightSet)
	fmt.Printf("merged %v\n", merged.Size())
	fmt.Printf("matrix %v\n", len(g.matrix)*len(g.matrix[0]))

	for _, row := range g.matrix {
		for _, tile := range row {
			// traverse close graf

			merged.Add(tile.Position)
		}
	}

	fmt.Printf("merged after %v\n", merged.Size())

	//setOfOthers := NewPositionSet()

	//for _, cols := range g.matrix {
	//	for _, tile := range cols {
	//		if !
	//		tile.Position
	//	}
	//}

	fmt.Printf("leftSetIsOpen: %v %d\n", leftSetIsOpen, leftSet.Size())
	fmt.Printf("rightSetIsOpen: %v %d\n", rightSetIsOpen, rightSet.Size())

	//fmt.Printf("leftSet %+v\n", leftSet)
	//fmt.Printf("rightSet %+v\n", rightSet)

	if leftSetIsOpen && !rightSetIsOpen {
		//fmt.Printf("leftSet %+v\n", leftSet)
		return rightSet.Size()
	}

	if rightSetIsOpen && !leftSetIsOpen {
		return leftSet.Size()
	}

	panic("one set should be closed and other open")
}

func (g *Graph) isOpenSet(s *PositionSet) bool {
	bh := g.boundaryHelper
	return s.minRow == bh.minRow || s.minCol == bh.minCol || s.maxCol == bh.maxCol || s.maxRow == bh.maxRow
}

func (g *Graph) traverseInWidth(set *PositionSet, pos Position) {
	if !g.boundaryHelper.Contains(pos) {
		return
	}

	g.matrix[pos.row][pos.col].TraverseInWidth(set)
}

type PositionSet struct {
	innerMap       map[Position]struct{}
	minRow, maxRow int
	minCol, maxCol int
}

func NewPositionSet() *PositionSet {
	return &PositionSet{
		innerMap: make(map[Position]struct{}),
		minRow:   math.MaxInt,
		maxRow:   math.MinInt,
		minCol:   math.MaxInt,
		maxCol:   math.MinInt,
	}
}

func (s *PositionSet) ForEach(fn func(position Position)) {
	for item, _ := range s.innerMap {
		fn(item)
	}
}

func (s *PositionSet) Clone() *PositionSet {
	clone := NewPositionSet()
	s.ForEach(clone.Add)
	return clone
}

func (s *PositionSet) Merge(others ...*PositionSet) *PositionSet {
	clone := s.Clone()

	for _, other := range others {
		other.ForEach(clone.Add)
	}

	return clone
}

func (s *PositionSet) Intersection(others ...*PositionSet) *PositionSet {
	merged := s.Merge(others...)
	inter := NewPositionSet()

	others = append(others, s)
	merged.ForEach(func(position Position) {
		for _, other := range others {
			if !other.Has(position) {
				return
			}
		}

		inter.Add(position)
	})

	return inter
}

func (s *PositionSet) Has(p Position) bool {
	_, wasVisited := s.innerMap[p]
	return wasVisited
}

func (s *PositionSet) Add(p Position) {
	s.innerMap[p] = struct{}{}
	s.minRow = tools.Min(s.minRow, p.row)
	s.maxRow = tools.Max(s.maxRow, p.row)
	s.minCol = tools.Min(s.minCol, p.col)
	s.maxCol = tools.Max(s.maxCol, p.col)
}

func (s *PositionSet) MinCorner() Position {
	return Position{
		row: s.minRow,
		col: s.minCol,
	}
}

func (s *PositionSet) MaxCorner() Position {
	return Position{
		row: s.maxRow,
		col: s.maxCol,
	}
}

func (s *PositionSet) Size() int {
	return len(s.innerMap)
}

//leftSetIsClosed: true 922
//rightSetIsClosed: false 258
