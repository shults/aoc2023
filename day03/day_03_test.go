package day03

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	buf := bytes.NewBuffer([]byte("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..\n"))
	p1, gearRatio := run(buf)
	assert.Equal(t, 4361, p1)          // part 1
	assert.Equal(t, 467835, gearRatio) // part 2
}

func TestMyTestCase(t *testing.T) {
	file, err := os.Open("03.in")

	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	res, gearRatio := run(file)

	assert.Equal(t, 559667, res)
	assert.Equal(t, 86841457, gearRatio)
}

func TestSymbolSet(t *testing.T) {
	ss := newSymbolBitSet(64)
	ss.markAsSymbol(0, 0)
	ss.markAsSymbol(0, 63)

	assert.True(t, ss.isSymbol(0, 0))
	assert.True(t, ss.isSymbol(0, 63))

	for col := 1; col < 63; col++ {
		assert.False(t, ss.isSymbol(0, col))
	}

	ss.markAsSymbol(1, 1)
	assert.True(t, ss.isSymbol(1, 1))
	for col := 0; col < 64; col++ {
		if col == 1 {
			continue
		}
		assert.False(t, ss.isSymbol(1, col))
	}

	ss.markAsSymbol(2, 2)
	assert.True(t, ss.isSymbol(2, 2))

	for col := 0; col < 64; col++ {
		if col == 2 {
			continue
		}
		assert.False(t, ss.isSymbol(2, col))
	}

	assert.False(t, ss.isSymbol(-1, 0))
	assert.False(t, ss.isSymbol(3, 0))
	assert.False(t, ss.isSymbol(0, 64))
	assert.False(t, ss.isSymbol(0, -1))
}

func TestNumberMachine(t *testing.T) {
	rowData := []byte("1..#...22.34")
	nm := newLineNumberMachine()
	nm2 := newLineNumberMachine()

	for coll, symbol := range rowData {
		nm.parseSymbol(symbol, 0, coll)
		nm2.parseSymbol(symbol, 2, coll)
	}

	for _, n := range []numberMachine{nm, nm2} {
		nums := n.getNumbers()
		assert.Equal(t, 3, len(nums))

		assert.Equal(t, 1, nums[0].val)
		assert.Equal(t, 22, nums[1].val)
		assert.Equal(t, 34, nums[2].val)
	}

	pairStructs := nm.getNumbers()[0].getNeighborPairs()

	pairs := make([]int, 0)

	for _, p := range pairStructs {
		pairs = append(pairs, p.row, p.coll)
	}

	assert.Equal(t, []int{
		-1, -1,
		-1, 0,
		-1, 1,
		0, -1,
		0, 1,
		1, -1,
		1, 0,
		1, 1,
	}, pairs)

	pairStructs = nm2.getNumbers()[0].getNeighborPairs()
	pairs = make([]int, 0, len(pairStructs))

	assert.True(t, len(pairStructs) == cap(pairStructs), "expected slice to be allocated with good capacity")

	for _, p := range pairStructs {
		pairs = append(pairs, p.row, p.coll)
	}

	assert.Equal(t, []int{
		1, -1,
		1, 0,
		1, 1,
		2, -1,
		2, 1,
		3, -1,
		3, 0,
		3, 1,
	}, pairs)
}

func TestDummy(t *testing.T) {
	gearMap := map[rowCollPair]struct{}{}

	gearMap[rowCollPair{coll: 1, row: 1}] = struct{}{}

	_, exists := gearMap[rowCollPair{coll: 1, row: 2}]

	assert.False(t, exists)

	_, exists = gearMap[rowCollPair{coll: 1, row: 1}]
	assert.True(t, exists)

	_, exists = gearMap[rowCollPair{coll: 1, row: 0}]
	assert.False(t, exists)
}
