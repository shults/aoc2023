package day06

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

func Main(flagSet *flag.FlagSet, args []string, in io.Reader) {
	parser := NewInputParser()

	start := time.Now()
	verbose := flagSet.Bool("verbose", false, "verbose mode")
	inputFile := flagSet.String("f", "", "input file")
	flagSet.BoolVar(&parser.force, "b", false, "use brut force method")
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

	part1, part2, calcTimeStart := run(
		in,
		*verbose,
		parser,
	)

	fmt.Printf("part1=%d\n", part1)
	fmt.Printf("part2=%d\n", part2)
	fmt.Printf("time microseconds all=%d\n", time.Since(start).Microseconds())
	fmt.Printf("time microseconds calc=%d\n", time.Since(calcTimeStart).Microseconds())
}

func run(in io.Reader, verbose bool, parser InputParser) (part1, part2 int, calcTimeStart time.Time) {
	reader := bufio.NewReader(in)

	var line []byte

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

	calcTimeStart = time.Now()

	part1 = parser.ProcessPart1(verbose)
	part2 = parser.ProcessPart2(verbose)

	return
}

func NewInputParser() InputParser {
	return InputParser{}
}

type sRace struct {
	time     int
	distance int
}

func (p *sRace) solutionsBrutForce() int {
	counter := 0

	for t := 0; t < p.time; t++ {
		speed := t
		travelTime := p.time - t
		distance := speed * travelTime

		if distance > p.distance {
			counter++
		}
	}

	return counter
}

func floor(x float64) float64 {
	if math.Floor(x) == x {
		return math.Floor(x - 1)
	} else {
		return math.Floor(x)
	}
}

func ceil(x float64) float64 {
	if math.Ceil(x) == x {
		return math.Ceil(x + 1)
	} else {
		return math.Ceil(x)
	}
}

func minMax(a, b float64) (float64, float64) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func (p *sRace) solutionsAnalytics() int {
	b := float64(p.time)
	c := float64(p.distance)
	d := b*b - 4*c

	if d <= 0 {
		return 0
	}

	dSqrt := math.Sqrt(d)

	x1, x2 := minMax((b-dSqrt)/2, (b+dSqrt)/2)

	return int(floor(x2)) - int(math.Max(ceil(x1), 0)) + 1
}

func (p *sRace) solutions(force bool) int {
	if force {
		return p.solutionsBrutForce()
	}
	return p.solutionsAnalytics()
}

func (p *sRace) consume(race sRace) {
	p.distance = concatInt(p.distance, race.distance)
	p.time = concatInt(p.time, race.time)
}

type InputParser struct {
	races []sRace
	// use brut force method
	force bool
}

func (p *InputParser) parseLine(line []byte) {
	sLine := string(line)

	if sTimes, ok := strings.CutPrefix(sLine, "Time:"); ok {
		rTimes := p.parseNumbers(sTimes)

		if p.races == nil {
			p.races = make([]sRace, len(rTimes))
		}

		for key, rTime := range rTimes {
			p.races[key].time = rTime
		}
	}

	if sDistances, ok := strings.CutPrefix(sLine, "Distance:"); ok {
		nDistances := p.parseNumbers(sDistances)

		if p.races == nil {
			p.races = make([]sRace, len(nDistances))
		}

		for key, distance := range nDistances {
			p.races[key].distance = distance
		}
	}
}

func (p *InputParser) parseNumbers(line string) []int {
	var res []int
	num := 0

	for _, val := range []byte(line) {
		if isAsciiDigit(val) {
			num *= 10
			num += int(val - '0')
			continue
		} else if num > 0 {
			res = append(res, num)
			num = 0
		}
	}

	if num > 0 {
		res = append(res, num)
		num = 0
	}

	return res
}

func isAsciiDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (p *InputParser) Print() {
	fmt.Printf("%+v\n", p)
}

func (p *InputParser) ProcessPart1(verbose bool) int {
	res := 1

	if verbose {
		fmt.Printf("ProcessPart1 ")
	}

	for _, race := range p.races {
		part := race.solutions(p.force)
		res *= part

		if verbose {
			fmt.Printf("%d ", part)
		}
	}

	if verbose {
		fmt.Printf("\n")
	}

	return res
}

func (p *InputParser) ProcessPart2(verbose bool) int {
	var masterRace sRace

	for _, childRace := range p.races {
		masterRace.consume(childRace)
	}

	if verbose {
		fmt.Printf("masterRace=%+v\n", masterRace)
	}

	return masterRace.solutions(p.force)
}

func concatInt(dest, src int) int {
	tmp := src
	multiplier := 1

	for {
		if tmp == 0 {
			break
		}

		tmp = tmp / 10
		multiplier *= 10
	}

	return multiplier*dest + src
}
