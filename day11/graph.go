package day11

import (
	"aoc2023/tools"
	"fmt"
)

type Universe struct {
	matrix   [][]*node
	galaxies []*node
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

	dimensions := tools.MatrixSize(matrix)

	for i := len(galaxiesInRows) - 1; i > -1; i-- {
		size := galaxiesInRows[i]

		if size == 0 {
			tools.MatrixInsertRow(&matrix, i, newNodes('.', dimensions.Cols))
			dimensions = tools.MatrixSize(matrix)
		}
	}

	for j := len(galaxiesInCols) - 1; j > -1; j-- {
		size := galaxiesInCols[j]

		if size == 0 {
			tools.MatrixInsertCol(&matrix, j, newNodes('.', dimensions.Rows))
			dimensions = tools.MatrixSize(matrix)
		}
	}

	var galaxies []*node

	for i, row := range matrix {
		for j, n := range row {
			n.setPosition(newPosition(i, j))

			if n.isGalaxy() {
				galaxies = append(galaxies, n)
			}
		}
	}

	return Universe{
		galaxies: galaxies,
		matrix:   matrix,
	}
}

func (u *Universe) Part1() int {
	printMatrix(u.matrix)

	positionPairs := newPositionPairSet()
	distanceSum := 0
	for i := 0; i < len(u.galaxies); i++ {
		galaxyA := u.galaxies[i]

		for j := i; j < len(u.galaxies); j++ {
			galaxyB := u.galaxies[j]

			if i == j {
				continue
			}

			if positionPairs.has(galaxyA.pos, galaxyB.pos) {
				continue
			}

			distance := galaxyA.distance(galaxyB)

			distanceSum += distance

			positionPairs.addPair(galaxyA.pos, galaxyB.pos, distance)

			fmt.Printf("%d to %d => %d\n", i+1, j+1, distance)
		}
	}

	return distanceSum
}

func (u *Universe) Part2() int {
	return -1
}

type node struct {
	symbol  byte
	related []*node
	pos     position
}

func newNode(symbol byte) *node {
	tools.AssertTrue(symbol == '#' || symbol == '.', "expected dot(.) or sharp(#)")

	return &node{
		symbol:  symbol,
		related: nil,
	}
}

func newNodes(symbol byte, size int) []*node {
	nodes := make([]*node, size)

	for i := 0; i < size; i++ {
		nodes[i] = newNode(symbol)
	}

	return nodes
}

func (n *node) isGalaxy() bool {
	return n.symbol == '#'
}

func (n *node) setPosition(pos position) {
	n.pos = pos
}

func (n *node) distance(other *node) int {
	return n.pos.distance(other.pos)
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

type positionPair struct {
	a position
	b position
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
