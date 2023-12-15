package day14

import (
	"strings"
)

type Plane struct {
	tiles [][]byte
}

func NewPlane(data []string) Plane {
	items := make([][]byte, len(data))

	for i, line := range data {
		items[i] = []byte(line)
	}

	return Plane{
		tiles: items,
	}
}

func (p *Plane) CalculatePart1(verbose bool) int {
	p.tilt()

	if verbose {
		p.String()
	}

	return p.calculateTotalLoad()
}

func (p *Plane) calculateTotalLoad() int {
	weight := len(p.tiles)
	sum := 0

	for i := 0; i < len(p.tiles); i++ {
		numberOfRoundStones := 0

		for j := 0; j < len(p.tiles[i]); j++ {
			if p.tiles[i][j] == 'O' {
				numberOfRoundStones++
			}
		}

		sum += weight * numberOfRoundStones
		weight--
	}

	return sum
}

func (p *Plane) CalculatePart2(verbose bool) int {
	cd := NewCycleDetector()

	for {
		p.doCycle()

		if cd.AddAndTryDetectCycle(p.calculateTotalLoad()) {
			break
		}
	}

	return cd.Predict(1_000_000_000 - 1)
}

func (p *Plane) tilt() {
	for i := 1; i < len(p.tiles); i++ {
		row := p.tiles[i]

		for j := 0; j < len(row); j++ {
			p.moveItem(i, j)
		}
	}
}

func (p *Plane) moveItem(i, j int) {
	if i == 0 {
		return
	}

	if p.tiles[i][j] == 'O' && p.tiles[i-1][j] == '.' {
		p.tiles[i][j], p.tiles[i-1][j] = p.tiles[i-1][j], p.tiles[i][j]
	}

	p.moveItem(i-1, j)
}

func (p *Plane) RotateRight() {
	n := len(p.tiles) - 1

	for i, maxi := 0, len(p.tiles); i < maxi; i++ {
		for j, maxj := 0, len(p.tiles[0])-i-1; j < maxj; j++ {
			p.swap(i, j, n-j, n-i)
		}
	}

	for i, maxi := 0, len(p.tiles)/2; i < maxi; i++ {
		for j := 0; j < len(p.tiles[0]); j++ {
			p.swap(i, j, n-i, j)
		}
	}
}

func (p *Plane) swap(ia, ja, ib, jb int) {
	p.tiles[ia][ja], p.tiles[ib][jb] = p.tiles[ib][jb], p.tiles[ia][ja]
}

func (p *Plane) doCycle() {
	for i := 0; i < 4; i++ {
		p.tilt()
		p.RotateRight()
	}
}

func (p *Plane) String() string {
	var sb strings.Builder

	for i := 0; i < len(p.tiles); i++ {
		for j := 0; j < len(p.tiles[i]); j++ {
			sb.WriteByte(p.tiles[i][j])
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
