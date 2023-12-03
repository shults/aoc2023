package day03

import (
	"bufio"
	"flag"
	"fmt"
	"io"
)

func Main(_ *flag.FlagSet, _ []string, in io.Reader) {
	multiplyResult, gearRatio := run(in)
	fmt.Printf("part1 = %d\npart2 = %d\n\n", multiplyResult, gearRatio)
}

func run(reader io.Reader) (sum int, gearRatioSum int) {
	br := bufio.NewReader(reader)
	row := 0
	sum = 0
	gearRatioSum = 0

	var ss *symbolSet
	var line []byte
	var nums []numEntity

	gearSetProduction := map[rowCollPair][]numEntity{}

	for {
		linePart, isPrefix, err := br.ReadLine()

		line = append(line, linePart...)

		if isPrefix {
			continue
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		if ss == nil {
			ss = newSymbolSet(len(line))
		}

		nm := newLineNumberMachine()

		for coll, symbol := range line {
			nm.parseSymbol(symbol, row, coll)

			if symbol == '*' {
				gearSetProduction[rowCollPair{row, coll}] = nil
			}

			if symbol == '.' {
				continue
			}

			ss.markAsSymbol(row, coll)
		}

		nums = append(nums, nm.getNumbers()...)

		line = nil
		row++
	}

numIteration:
	for _, num := range nums {

		pairs := num.getNeighborPairs()

		for _, pair := range pairs {
			items, isGear := gearSetProduction[pair]

			if isGear {
				gearSetProduction[pair] = append(items, num)
			}
		}

	neighboursIteration:
		for _, pair := range pairs {

			if !ss.isSymbol(pair.row, pair.coll) {
				continue neighboursIteration
			}

			sum += num.val
			continue numIteration
		}
	}

	for _, elems := range gearSetProduction {
		if len(elems) < 2 {
			continue
		}

		production := 1

		for _, elem := range elems {
			production *= elem.val
		}

		gearRatioSum += production
	}

	return
}

func newSymbolSet(width int) *symbolSet {
	return &symbolSet{
		data:  make([][]uint64, 0),
		width: width,
	}
}

type symbolSet struct {
	data  [][]uint64
	width int
}

func (ss *symbolSet) markAsSymbol(row, coll int) {
	if row < 0 || coll < 0 {
		return
	}

	for i := len(ss.data); i <= row; i++ {
		ss.data = append(ss.data, make([]uint64, (ss.width/64)+1))
	}

	rowData := ss.data[row]

	rowData[coll/64] = rowData[coll/64] | (1 << (coll % 64))
}

func (ss *symbolSet) isSymbol(row, coll int) bool {
	// out of boundary
	if row < 0 || coll < 0 || coll > ss.width-1 || row > len(ss.data)-1 {
		return false
	}

	rowData := ss.data[row]

	return rowData[coll/64]&(1<<(coll%64)) > 0
}

func isAsciiDigit(symbol byte) bool {
	return symbol >= '0' && symbol <= '9'
}

type numEntity struct {
	val      int
	row      int
	colStart int
	colEnd   int
}

type rowCollPair struct {
	row, coll int
}

func (ne *numEntity) getNeighborPairs() []rowCollPair {
	pairs := make([]rowCollPair, 0, 2+2*(ne.colEnd-ne.colStart+2))

	for i := ne.row - 1; i < ne.row+2; i++ {
		for j := ne.colStart - 1; j < ne.colEnd+1; j++ {
			if i == ne.row && j >= ne.colStart && j < ne.colEnd {
				continue
			}

			pairs = append(pairs, rowCollPair{
				row:  i,
				coll: j,
			})
		}
	}

	return pairs
}

func newLineNumberMachine() numberMachine {
	return numberMachine{
		numbers: make([]numEntity, 0),
		state:   &symbolStateMachineParser{},
	}
}

type numberMachine struct {
	numbers []numEntity
	state   stateMachineParser
}

func (nm *numberMachine) parseSymbol(symbol byte, row int, coll int) {
	num, nextState := nm.state.parse(symbol, row, coll)

	if num != nil {
		nm.numbers = append(nm.numbers, *num)
	}

	nm.state = nextState
}

func (nm *numberMachine) getNumbers() []numEntity {
	num := nm.state.flush()
	res := make([]numEntity, 0, len(nm.numbers)+1)
	res = append(res, nm.numbers...)

	if num != nil {
		res = append(res, *num)
	}

	return res
}

type stateMachineParser interface {
	parse(symbol byte, row int, coll int) (*numEntity, stateMachineParser)
	flush() *numEntity
}

func newSymbolStateMachineParser() stateMachineParser {
	return &symbolStateMachineParser{}
}

type symbolStateMachineParser struct{}

func (s *symbolStateMachineParser) flush() *numEntity {
	return nil
}

func (s *symbolStateMachineParser) parse(symbol byte, row int, coll int) (*numEntity, stateMachineParser) {
	if isAsciiDigit(symbol) {
		return nil, newNumberStateMachineParse(int(symbol-'0'), row, coll)
	} else {
		return nil, newSymbolStateMachineParser()
	}
}

func newNumberStateMachineParse(val, row, coll int) *numberStateMachineParse {
	return &numberStateMachineParse{
		num: numEntity{
			val:      val,
			colStart: coll,
			colEnd:   coll + 1,
			row:      row,
		},
	}
}

type numberStateMachineParse struct {
	num numEntity
}

func (n *numberStateMachineParse) flush() *numEntity {
	return &n.num
}

func (n *numberStateMachineParse) parse(symbol byte, row int, coll int) (*numEntity, stateMachineParser) {
	if isAsciiDigit(symbol) {
		n.num.val *= 10
		n.num.val += int(symbol - '0')
		n.num.colEnd = coll + 1
		return nil, n
	}

	return &n.num, &symbolStateMachineParser{}
}
