package day14

import (
	"fmt"
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
	p.tiltNorth()

	if verbose {
		p.print()
	}

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
	return -1
}

func (p *Plane) tiltNorth() {
	for i := 1; i < len(p.tiles); i++ {
		row := p.tiles[i]

		for j := 0; j < len(row); j++ {
			p.moveItemNorth(i, j)
		}
	}
}

func (p *Plane) moveItemNorth(i, j int) {
	if i == 0 {
		return
	}

	if p.tiles[i][j] == 'O' && p.tiles[i-1][j] == '.' {
		p.tiles[i][j], p.tiles[i-1][j] = p.tiles[i-1][j], p.tiles[i][j]
		p.moveItemNorth(i-1, j)
	}
}

func (p *Plane) print() {
	var sb strings.Builder

	for i := 0; i < len(p.tiles); i++ {
		for j := 0; j < len(p.tiles[i]); j++ {
			sb.WriteByte(p.tiles[i][j])
		}
		sb.WriteByte('\n')
	}

	fmt.Printf("%s\n", sb.String())
}
