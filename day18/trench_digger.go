package day18

import (
	"aoc2023/tools"
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type trenchDigger struct {
	instructions []instruction
}

func newTrenchDigger(rawInstructions []string) trenchDigger {
	instructions := make([]instruction, len(rawInstructions))

	for i, ins := range rawInstructions {
		instructions[i] = newInstruction(ins)
	}

	return trenchDigger{
		instructions: instructions,
	}
}

func (t *trenchDigger) square(verbose bool) int {
	tm := newTrenchMap(t.instructions)
	return tm.calculateSquare(verbose)
}

type instruction struct {
	direction
	length int
}

func newInstruction(ins string) instruction {
	fields := strings.Fields(ins)
	tools.AssertTrue(len(fields) == 3)
	length, err := strconv.Atoi(fields[1])
	tools.AssertNoError(err)

	return instruction{
		direction: newDirection(fields[0]),
		length:    length,
	}
}

var allowedDirections = map[byte]struct{}{
	'U': {},
	'D': {},
	'L': {},
	'R': {},
}

type direction struct {
	dir byte
}

func (d *direction) is(b byte) bool {
	return d.dir == b
}

func newDirection(s string) direction {
	bytes := []byte(s)
	tools.AssertTrue(len(bytes) == 1)
	_, ok := allowedDirections[bytes[0]]
	tools.AssertTrue(ok, "non expected direction")

	return direction{
		dir: bytes[0],
	}
}

type position struct {
	i, j int
}

func (p position) up() position {
	p.i--
	return p
}

func (p position) down() position {
	p.i++
	return p
}

func (p position) left() position {
	p.j--
	return p
}

func (p position) right() position {
	p.j++
	return p
}

type node struct {
	position
	neighbours []*node
}

func (n *node) bind(other *node) {
	n.neighbours = append(n.neighbours, other)
	other.neighbours = append(other.neighbours, n)
}

type trenchMap struct {
	trenchData map[position]*node
}

func newTrenchMap(instructions []instruction) *trenchMap {
	tm := make(map[position]*node)
	start := &node{position{0, 0}, nil}
	tm[start.position] = start

	current := start
	for _, ins := range instructions {
		for ins.length > 0 {
			var pos position

			switch ins.dir {
			case 'U':
				pos = current.up()
			case 'D':
				pos = current.down()
			case 'L':
				pos = current.left()
			case 'R':
				pos = current.right()
			default:
				panic("non expected direction")
			}

			var nextNode *node

			if n, ok := tm[pos]; ok {
				nextNode = n
				nextNode.neighbours = append(nextNode.neighbours, current)
			} else {
				nextNode = &node{
					position: pos,
					neighbours: []*node{
						current,
					},
				}
				tm[pos] = nextNode
			}

			current.neighbours = append(current.neighbours, nextNode)
			current = nextNode

			ins.length--
		}
	}

	for _, n := range tm {
		tools.AssertTrue(len(n.neighbours) == 2)
	}

	return &trenchMap{
		trenchData: tm,
	}
}

func (t *trenchMap) getCorners() (min, max position) {
	min = position{0, 0}
	max = position{0, 0}

	for pos, n := range t.trenchData {
		tools.AssertTrue(len(n.neighbours) == 2)

		min.i = tools.Min(min.i, pos.i)
		min.j = tools.Min(min.j, pos.j)

		max.i = tools.Max(max.i, pos.i)
		max.j = tools.Max(max.j, pos.j)
	}

	return
}

func (t *trenchMap) perimeter() []position {
	min, max := t.getCorners()

	var perItems []position

	for i := min.i; i <= max.i; i++ {
		for j := min.j; j <= max.j; j++ {
			if i > min.i && i < max.i && j > min.j && j < max.j {
				continue
			}

			perItems = append(perItems, position{i, j})
		}
	}

	return perItems
}

func (t *trenchMap) calculateSquare(verbose bool) int {
	min, max := t.getCorners()
	grid := makeGrid(min, max)

	for i := min.i; i <= max.i; i++ {
		for j := min.j; j <= max.j; j++ {
			pos := position{i, j}

			if _, isTrench := t.trenchData[pos]; isTrench {
				continue
			}

			n := grid.Get(pos)

			for _, nPos := range []position{
				pos.up(),
				pos.down(),
				pos.left(),
				pos.right(),
			} {
				if _, isTrench := t.trenchData[nPos]; isTrench {
					continue
				}

				if !grid.IsInside(nPos) {
					continue
				}

				grid.Get(nPos).bind(n)
			}
		}
	}

	outOfTrench := make(map[position]struct{})

	for _, pos := range t.perimeter() {
		nodes := []*node{
			grid.Get(pos),
		}

		for {
			if len(nodes) == 0 {
				break
			}

			n := nodes[0]
			nodes = nodes[1:]

			if _, isTrench := t.trenchData[n.position]; isTrench {
				continue
			}

			if _, wasVisited := outOfTrench[n.position]; wasVisited {
				continue
			}

			outOfTrench[n.position] = struct{}{}

			nodes = append(nodes, n.neighbours...)
		}
	}

	if verbose {
		t.print(outOfTrench)
	}

	width := max.j - min.j + 1
	height := max.i - min.i + 1

	return width*height - len(outOfTrench)
}

func (t *trenchMap) print(outOfTrench map[position]struct{}) {
	leftCornet, rightCorner := t.getCorners()
	yellow := color.New(color.BgYellow)

	var sb strings.Builder
	for i := leftCornet.i; i <= rightCorner.i; i++ {
		for j := leftCornet.j; j <= rightCorner.j; j++ {
			pos := position{i, j}

			if _, ok := t.trenchData[pos]; ok {
				sb.WriteByte('#')
			} else if _, ok := outOfTrench[pos]; ok {
				sb.WriteString(yellow.Sprintf("%s", "*"))
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}

	fmt.Printf("%s\n", sb.String())
}

func makeGrid(min, max position) grid {
	height := max.i - min.i + 1
	width := max.j - min.j + 1

	return grid{
		data:  make([]*node, width*height),
		min:   min,
		max:   max,
		width: width,
	}
}

type grid struct {
	data     []*node
	min, max position
	width    int
}

func (g *grid) IsInside(p position) bool {
	return p.i >= g.min.i && p.j >= g.min.j && p.i <= g.max.i && p.j <= g.max.j
}

func (g *grid) Get(p position) *node {
	offset := g.offset(p)

	n := g.data[offset]

	if n == nil {
		n = &node{
			position: p,
		}
		g.data[offset] = n
	}

	return n
}

func (g *grid) Set(p position, n *node) {
	offset := g.offset(p)
	g.data[offset] = n
}

func (g *grid) offset(p position) int {
	return (p.i-g.min.i)*g.width + (p.j - g.min.j)
}
