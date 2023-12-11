package day11

import (
	"aoc2023/tools"
	"fmt"
)

type Universe struct {
	matrix [][]*node
}

func NewUniverse(data []string) Universe {
	galaxiesInRows := make([]int, len(data))
	galaxiesInCols := make([]int, len(data[0]))

	matrix := make([][]*node, len(data))

	for i, row := range data {
		matrix[i] = make([]*node, len(data))

		for j, symbol := range []byte(row) {
			n := newNode(symbol)

			if n.isGalaxy() {
				galaxiesInRows[i]++
				galaxiesInCols[j]++
			}

			matrix[i][j] = n
		}
	}

	return Universe{
		matrix: matrix,
	}
}

func (u *Universe) getGalaxies(emptyRowDistance int) []*node {
	var galaxies []*node

	for i, row := range u.matrix {
		for j, n := range row {
			n.setPosition(newPosition(i, j))
			n.bindWithNeighbours(u.matrix)

			if n.isGalaxy() {
				galaxies = append(galaxies, n)
			}
		}
	}

	dim := tools.MatrixSize(u.matrix)

	for i := 0; i < dim.Rows; i++ {
		hasNoGalaxies := true

		for j := 0; j < dim.Cols; j++ {
			nodeItem := u.matrix[i][j]
			if nodeItem.isGalaxy() {
				hasNoGalaxies = false
			}
		}

		if hasNoGalaxies {
			for j := 0; j < dim.Cols; j++ {
				nodeItem := u.matrix[i][j]
				nodeItem.size = emptyRowDistance
			}
		}
	}

	for j := 0; j < dim.Cols; j++ {
		hasNoGalaxies := true

		for i := 0; i < dim.Rows; i++ {
			nodeItem := u.matrix[i][j]

			if nodeItem.isGalaxy() {
				hasNoGalaxies = false
			}
		}

		if hasNoGalaxies {
			for i := 0; i < dim.Cols; i++ {
				nodeItem := u.matrix[i][j]
				nodeItem.size = emptyRowDistance
			}
		}
	}

	return galaxies
}

func (u *Universe) traverseGalaxy(n *node, cb func(int, *node)) {
	type nodeDistancePair struct {
		distance int
		nPtr     *node
	}

	pairs := []nodeDistancePair{{
		nPtr:     n,
		distance: 0,
	}}

	visitList := make(map[*node]struct{})

	for {
		newPairs := make([]nodeDistancePair, 0)

		for _, pair := range pairs {
			if _, wasVisited := visitList[pair.nPtr]; wasVisited {
				continue
			}

			// mark as visited
			visitList[pair.nPtr] = struct{}{}

			if pair.nPtr.isGalaxy() {
				cb(pair.distance, pair.nPtr)
			}

			for _, neighbour := range pair.nPtr.neighbours {
				newPairs = append(newPairs, nodeDistancePair{
					nPtr:     neighbour,
					distance: pair.distance + neighbour.size,
				})
			}
		}

		pairs = newPairs

		// no pairs to visit
		if len(pairs) == 0 {
			break
		}
	}

}

func (u *Universe) calculate(emptyRowDistance int) int {
	galaxies := u.getGalaxies(emptyRowDistance)

	positionPairs := newPositionPairSet()
	distanceSum := 0

	for i := 0; i < len(galaxies)-1; i++ {
		galaxyA := galaxies[i]
		u.traverseGalaxy(galaxyA, func(distance int, galaxyB *node) {
			if positionPairs.has(galaxyA.pos, galaxyB.pos) {
				return
			}

			positionPairs.addPair(galaxyA.pos, galaxyB.pos, distance)
			distanceSum += distance
		})
	}

	return distanceSum
}

func (u *Universe) Part1() int {
	return u.calculate(2)
}

func (u *Universe) Part2(emptyRowOrColumnSize int) int {
	return u.calculate(emptyRowOrColumnSize)
}

type node struct {
	symbol     byte
	neighbours []*node
	pos        position
	size       int
}

func newNode(symbol byte) *node {
	tools.AssertTrue(symbol == '#' || symbol == '.', "expected dot(.) or sharp(#)")

	return &node{
		symbol:     symbol,
		neighbours: nil,
		size:       1,
	}
}

func (n *node) isGalaxy() bool {
	return n.symbol == '#'
}

func (n *node) setPosition(pos position) {
	n.pos = pos
}

func (n *node) shiftRow(dx int) {
	n.pos.row += dx
}

func (n *node) distance(other *node) int {
	return n.pos.distance(other.pos)
}

func (n *node) bindWithNeighbours(matrix [][]*node) {
	dim := tools.MatrixSize(matrix)
	neighbours := make([]*node, 0, 4)

	for _, loc := range []position{
		n.pos.up(),
		n.pos.down(),
		n.pos.left(),
		n.pos.right(),
	} {
		if !dim.Contains(loc.row, loc.col) {
			continue
		}

		neighbour := matrix[loc.row][loc.col]
		neighbours = append(neighbours, neighbour)
	}

	n.neighbours = neighbours
}

type position struct {
	row, col int
}

func newPosition(col, row int) position {
	return position{
		col, row,
	}
}

func (p position) distance(other position) int {
	return abs(p.row-other.row) + abs(p.col-other.col)
}

func (p position) up() position {
	p.row--
	return p
}

func (p position) down() position {
	p.row++
	return p
}

func (p position) left() position {
	p.col--
	return p
}

func (p position) right() position {
	p.col++
	return p
}

func newPositionPairSet() *positionPairSet {
	return &positionPairSet{
		pairMap: make(map[position]map[position]int),
	}
}

type positionPairSet struct {
	pairMap map[position]map[position]int
}

func (s *positionPairSet) addPair(a position, b position, distance int) {
	if a == b {
		return
	}

	s.addPairInner(a, b, distance)
	s.addPairInner(b, a, distance)
}

func (s *positionPairSet) addPairInner(a position, b position, distance int) {
	aMap, ok := s.pairMap[a]

	if !ok {
		aMap = make(map[position]int)
		s.pairMap[a] = aMap
	}

	aMap[b] = distance
}

func (s *positionPairSet) has(a position, b position) bool {
	if a == b {
		return false
	}

	aMap, has := s.pairMap[a]

	if !has {
		return false
	}

	_, hasB := aMap[b]

	return hasB
}

func printMatrix(matrix [][]*node) {
	for _, row := range matrix {
		for _, item := range row {
			fmt.Printf("%s", []byte{item.symbol})
		}
		fmt.Printf("\n")
	}
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}
