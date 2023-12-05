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
		for _, sub := range mapper.subMappers {
			fmt.Printf("\t%d %d %d\n", sub.dest.start, sub.src.start, sub.src.length())
		}
		fmt.Printf("\n")
	}
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

	var ranges []Range

	for i := 0; i < len(p.seeds); i += 2 {
		ranges = append(ranges, Range{
			start: p.seeds[i],
			end:   p.seeds[i] + p.seeds[i+1] - 1,
		})
	}

	p.processSeedPack(ranges, verbose)

	return lowesLocation
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

func (p *InputParser) processSeedPack(ranges []Range, verbose bool) []Range {
	for _, mapper := range p.mappers {
		ranges = mapper.MapRanges(ranges, verbose)
	}

	return ranges
}

func NewSeedMapper(name string) *DestMapper {
	return &DestMapper{
		name: name,
	}
}

type DestMapper struct {
	name       string
	subMappers []SubMapper
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

	m.subMappers = append(m.subMappers, SubMapper{
		src: Range{
			start: nums[1],
			end:   nums[1] + nums[2],
		},
		dest: Range{
			start: nums[0],
			end:   nums[0] + nums[2],
		},
	})
}

func (m *DestMapper) Map(src int) int {
	for _, subMapper := range m.subMappers {
		if subMapper.accepts(src) {
			return subMapper.convert(src)
		}
	}

	return src
}

func (m *DestMapper) MapRanges(ranges []Range, verbose bool) []Range {
	var res []Range

	for _, sr := range ranges {
		for _, ss := range m.subMappers {
			res = append(res, ss.splitRange(sr)...)
		}
	}

	return res
}

type SubMapper struct {
	src  Range
	dest Range
}

func (s *SubMapper) convert(src int) int {
	return s.dest.start + src - s.src.start
}

func (s *SubMapper) accepts(src int) bool {
	return s.src.contains(src)
}

func (s *SubMapper) splitRange(sr Range) []Range {

	return nil

	// todo: add case when data set is subset of range and split range into 3 parts

	//if sr.start >= s.sourceRangeStart && sr.end <= s.sourceEnd() {
	//	// todo: includes and split into three parts
	//	return []Range{
	//		{
	//			start: s.destinationRangeStart,
	//		},
	//	}
	//} else if sr.end >= s.sourceRangeStart && sr.end <= s.sourceEnd() {
	//	// todo: make left overlap + map
	//	return nil
	//} else if sr.start >= s.sourceRangeStart && sr.end <= s.sourceEnd() {
	//	// right overlap
	//	return nil
	//} else {
	//	// no overlap
	//	// dont split range
	//	return []Range{
	//		sr,
	//	}
	//}
}

type Range struct {
	start int
	end   int
}

func (r *Range) contains(src int) bool {
	return src >= r.start && src <= r.end
}

func (r *Range) length() int {
	return r.end - r.end
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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
