package day04

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

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

	points, scratchcards := 0, 0
	reader := bufio.NewReader(in)

	var line []byte
	var cards []*Card

	for {
		subLine, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("Got error: %s\n", err)
			os.Exit(1)
		}

		line = append(line, subLine...)

		if isPrefix {
			continue
		}

		card, err := ParseCard(line)

		if err != nil {
			panic(err)
		}

		cards = append(cards, &card)
		points += card.points

		line = nil
	}

	for i := 0; i < len(cards); i++ {
		card := cards[i]

		for j, top := i+1, min(len(cards), i+card.matches+1); j < top; j++ {
			cards[j].copies++
			cards[j].copies += card.copies
		}

		scratchcards += 1 + card.copies
	}

	if *verbose {
		for _, card := range cards {
			fmt.Printf("card = %+v\n", card)
		}
	}

	fmt.Printf("part1 = %d\nscratchcards = %d\n", points, scratchcards)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func ParseCard(line []byte) (Card, error) {
	id := 0
	pos := 0

	for {
		symbol := line[pos]
		pos++

		if symbol == ':' {
			break
		}

		if !isAsciiDigit(symbol) {
			continue
		}

		id *= 10
		id += int(symbol - '0')
	}

	num := 0

	winNums := make([]int, 0)

	for {
		symbol := line[pos]
		pos++

		if symbol == '|' {
			break
		}

		if !isAsciiDigit(symbol) {
			if num > 0 {
				winNums = append(winNums, num)
				num = 0
			}
			continue
		}

		num *= 10
		num += int(symbol - '0')
	}

	num = 0
	seqNums := make([]int, 0)
	for {
		if pos == len(line) {
			if num > 0 {
				seqNums = append(seqNums, num)
				num = 0
			}

			break
		}

		symbol := line[pos]
		pos++

		if !isAsciiDigit(symbol) {
			if num > 0 {
				seqNums = append(seqNums, num)
				num = 0
			}
			continue
		}

		num *= 10
		num += int(symbol - '0')
	}

	points, matches := getPointsAndMatches(winNums, seqNums)

	return Card{id: id, points: points, matches: matches, copies: 0}, nil
}

type Card struct {
	id      int
	points  int
	matches int
	copies  int
}

func getPointsAndMatches(winNums []int, seqNums []int) (points, matches int) {
	winSet := make(map[int]struct{}, len(winNums))

	for _, winNum := range winNums {
		winSet[winNum] = struct{}{}
	}

	matches = 0
	for _, seqNum := range seqNums {
		if _, contains := winSet[seqNum]; contains {
			matches += 1
		}
	}

	if matches == 0 {
		return 0, matches
	} else {
		return 1 << (matches - 1), matches
	}

}

func isAsciiDigit(symbol byte) bool {
	return symbol >= '0' && symbol <= '9'
}

//Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53 | 8 4 c
//Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19 | 2 2 c 1
//Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1 | 2 2 c 1 + 1 + 1(from copy) + original = total 4 instances
//Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83 | 1 1 c 1 + 1 + 1(from copy) + original = total 4 instances
//Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36 | 0 0 c 1(from 1) + 4(from 3) + 4(from 4) + original
//Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11 | 0 0
