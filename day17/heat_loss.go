package day17

import (
	"aoc2023/tools"
)

type HeatLossMatrix struct {
	tiles [][]int
	dim   tools.Dimensions
}

func NewHeatLossMatrix(rows []string) *HeatLossMatrix {
	tiles := make([][]int, len(rows))

	for i, row := range rows {
		tiles[i] = make([]int, len([]byte(row)))
		for j, symbol := range []byte(row) {
			tiles[i][j] = tools.AsciiNumberToInt(symbol)
		}
	}

	return &HeatLossMatrix{
		tiles: tiles,
		dim:   tools.MatrixSize(tiles),
	}
}

func (m *HeatLossMatrix) GetMinimalHeatLoss() int {
	dest := Position{0, 0}

	m.nextStep(dest.left(), NewVisitList().Add(dest))
	m.nextStep(dest.bottom(), NewVisitList().Add(dest))

	return -1
}

func (m *HeatLossMatrix) nextStep(
	pos Position,
	visitList VisitList,
) {
	if !m.isInside(pos) {
		return
	}

	if visitList.SameDirectionLimitReached() {
		return
	}

	if visitList.WasVisited(pos) {
		return
	}

	visitList = visitList.Add(pos)

	for _, nextPos := range pos.neighbours() {
		m.nextStep(nextPos, visitList)
	}
}

func (m *HeatLossMatrix) isInside(pos Position) bool {
	return pos.i >= 0 && pos.j >= 0 && pos.i < m.dim.Rows-1 && pos.j <= m.dim.Cols-1
}

type VisitList struct {
	head *VisitListNode
}

func NewVisitList() VisitList {
	return VisitList{}
}

func (v VisitList) Add(pos Position) VisitList {
	return VisitList{
		head: &VisitListNode{
			pos:  pos,
			prev: v.head,
		},
	}
}

func (v VisitList) WasVisited(pos Position) bool {
	return v.head.WasVisited(pos)
}

func (v VisitList) SameDirectionLimitReached() bool {
	limit := 3

	current := v.head
	directions := make([]Direction, limit+1)

	for i := 0; i < limit+1; i++ {
		if current == nil {
			return false
		}

		prev := current.prev

		if prev == nil {
			return false
		}

		directions[i] = prev.direction(current)

		current = prev
	}

	for i := 1; i < len(directions); i++ {
		if directions[i] != directions[0] {
			return false
		}
	}

	return true
}

type VisitListNode struct {
	pos  Position
	prev *VisitListNode
}

func (v *VisitListNode) WasVisited(pos Position) bool {
	if v == nil {
		return false
	}

	return v.pos == pos || v.prev.WasVisited(pos)
}

func (v *VisitListNode) direction(current *VisitListNode) Direction {
	return v.pos.direction(current.pos)
}
