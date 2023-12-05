package day05

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// part1: 251346198
// part2: 72263011

func Main(flagSet *flag.FlagSet, args []string, in io.Reader) {
	verbose := flagSet.Bool("verbose", false, "verbose mode")
	inputFile := flagSet.String("f", "", "input file")
	err := flagSet.Parse(args)

	if err != nil {
		flagSet.Usage()
		fmt.Printf("Error ocurred: %s\n", err)
		os.Exit(1)
	}

	if len(*inputFile) > 0 {
		in, err = os.Open(*inputFile)

		if err != nil {
			flagSet.Usage()
			fmt.Printf("Error ocurred: %s\n", err)
			os.Exit(1)
		}
	}

	part1, part2 := run(
		in,
		*verbose,
	)

	fmt.Printf("part1=%d\n", part1)
	fmt.Printf("part2=%d\n", part2)
}

func run(in io.Reader, verbose bool) (part1, part2 int) {
	reader := bufio.NewReader(in)

	var line []byte
	parser := NewInputParser()

	for {
		data, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		line = append(line, data...)

		if isPrefix {
			continue
		}

		parser.parseLine(line)

		line = line[:0]
	}

	if verbose {
		parser.Print()
	}

	part1 = parser.ProcessPart1(verbose)
	part2 = parser.ProcessPart2(verbose)

	return
}

func NewInputParser() InputParser {
	return InputParser{
		seeds: nil,
	}
}

type InputParser struct {
	seeds         []int
	mappers       []*DestMapper
	currentMapper *DestMapper
}

func (p *InputParser) parseLine(line []byte) {
	if len(line) == 0 {
		return
	}

	if p.seeds == nil && string(line[:6]) == "seeds:" {
		p.seeds = make([]int, 0)
		num := 0

		for _, symbol := range line[6:] {
			if isAsciiNumber(symbol) {
				num *= 10
				num += int(symbol - '0')
			} else if num > 0 {
				p.seeds = append(p.seeds, num)
				num = 0
			}
		}

		p.seeds = append(p.seeds, num)
		return
	}

	if line[len(line)-1] == ':' {
		p.currentMapper = NewSeedMapper(string(line[:len(line)-1]))
		p.mappers = append(p.mappers, p.currentMapper)
		return
	}

	p.currentMapper.ParseLine(line)
}

func (p *InputParser) Print() {
	fmt.Printf("seeds: %+v\n", p.seeds)
	fmt.Printf("maps:\n")

	for _, mapper := range p.mappers {
		fmt.Printf("\t%s:\n", mapper.name)
		for _, sub := range mapper.subSets {
			fmt.Printf("\t%d %d %d\n", sub.destinationRangeStart, sub.sourceRangeStart, sub.rangeLength)
		}
		fmt.Printf("\n")
	}
}

func (p *InputParser) processSeed(seed int, verbose bool) int {
	dest := seed

	if verbose {
		fmt.Printf("[%d", dest)
	}

	for _, mapper := range p.mappers {
		dest = mapper.Map(dest)

		if verbose {
			fmt.Printf(" => %d", dest)
		}
	}

	if verbose {
		fmt.Printf("]\n")
	}

	return dest
}

func (p *InputParser) ProcessPart1(verbose bool) int {
	lowesLocation := math.MaxInt

	for _, seed := range p.seeds {
		lowesLocation = min(lowesLocation, p.processSeed(seed, verbose))
	}

	return lowesLocation
}

func (p *InputParser) ProcessPart2(verbose bool) int {
	lowesLocation := math.MaxInt
	assert(len(p.seeds)%2 == 0, "expected to be even")

	for i := 0; i < len(p.seeds); i += 2 {
		for seed, max := p.seeds[i], p.seeds[i]+p.seeds[i+1]-1; seed < max; seed++ {
			lowesLocation = min(lowesLocation, p.processSeed(seed, verbose))
		}
	}

	return lowesLocation
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func NewSeedMapper(name string) *DestMapper {
	return &DestMapper{
		name: name,
	}
}

type DestMapper struct {
	name    string
	subSets []DestSet
}

func (m *DestMapper) ParseLine(line []byte) {
	numStr := strings.Split(string(line), " ")
	assert(len(numStr) == 3, "expected 3")
	var nums []int

	for _, val := range numStr {
		num, err := strconv.Atoi(val)
		assertNoError(err)
		nums = append(nums, num)
	}

	m.subSets = append(m.subSets, DestSet{
		destinationRangeStart: nums[0],
		sourceRangeStart:      nums[1],
		rangeLength:           nums[2],
	})
}

func (m *DestMapper) Map(src int) (dest int) {
	dest = src

	for _, subSet := range m.subSets {
		if dest >= subSet.sourceRangeStart && dest <= subSet.sourceRangeStart+subSet.rangeLength-1 {
			dest = subSet.destinationRangeStart + dest - subSet.sourceRangeStart
			return
		}
	}

	return
}

type DestSet struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func isAsciiNumber(symbol byte) bool {
	return symbol >= '0' && symbol <= '9'
}

func assert(assertion bool, msg string) {
	if !assertion {
		panic(msg)
	}
}

func assertNoError(err error) {
	if err != nil {
		panic(err)
	}
}
