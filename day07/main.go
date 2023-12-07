package day07

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type CardCombinationScore int8

const (
	// HighCard where all cards' labels are distinct: 23456
	HighCard CardCombinationScore = iota

	// OnePair where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	OnePair

	// TwoPair where two cards share one label, two other cards share a second label, and the remaining CardKind has a third label: 23432
	TwoPair

	// ThreeOfAKind where three cards have the same label, and the remaining two cards are each different from any other CardKind in the hand: TTT98
	ThreeOfAKind

	// FullHouse where three cards have the same label, and the remaining two cards share a different label: 23332
	FullHouse

	//	FourOfAKind where four cards have the same label and one CardKind has a different label: AA8AA
	FourOfAKind

	// FiveOfAKind where all five cards have the same label: AAAAA
	FiveOfAKind
)

type CardKind = int8

const (
	CardA CardKind = 12 - iota
	CardK
	CardQ
	CardT
	Card9
	Card8
	Card7
	Card6
	Card5
	Card4
	Card3
	Card2
	CardJ
)

var (
	cardMap = map[byte]CardKind{
		'A': CardA,
		'K': CardK,
		'Q': CardQ,
		'J': CardJ,
		'T': CardT,
		'9': Card9,
		'8': Card8,
		'7': Card7,
		'6': Card6,
		'5': Card5,
		'4': Card4,
		'3': Card3,
		'2': Card2,
	}
)

func Main(flagSet *flag.FlagSet, args []string, in io.Reader) {
	start := time.Now()
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

	part1, part2, calcTimeStart := run(
		in,
		*verbose,
	)

	fmt.Printf("part1=%d\n", part1)
	fmt.Printf("part2=%d\n", part2)
	fmt.Printf("time microseconds all=%d\n", time.Since(start).Microseconds())
	fmt.Printf("time microseconds calc=%d\n", time.Since(calcTimeStart).Microseconds())
}

func run(in io.Reader, verbose bool) (part1, part2 int, calcTimeStart time.Time) {
	reader := bufio.NewReader(in)
	cf := newCardFactory()

	var combinations []CardsCombination

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		assertNoError(err)

		combinations = append(combinations, cf.parseCombination(string(line)))
	}

	calcTimeStart = time.Now()

	slices.SortFunc(combinations, func(a, b CardsCombination) int {
		return a.Compare(&b)
	})

	part1 = 0
	for i, cc := range combinations {
		part1 += (i + 1) * cc.score
	}

	if verbose {
		fmt.Printf("%+v\n", combinations)
	}

	return
}

func newCardFactory() cardFactory {
	return cardFactory{}
}

type cardFactory struct {
	combinations [13]uint8
}

func (c *cardFactory) getCombination(cards *FiveCards) CardCombinationScore {
	for i := 0; i < len(c.combinations); i++ {
		c.combinations[i] = 0
	}

	for _, card := range cards {
		c.combinations[card] += 1
	}

	threeOfAkind := false
	pairsNumber := 0

	for _, card := range cards {
		if c.combinations[card] == 5 {
			return FiveOfAKind
		} else if c.combinations[card] == 4 {
			return FourOfAKind
		} else if c.combinations[card] == 3 {
			c.combinations[card] = 0
			threeOfAkind = true
		} else if c.combinations[card] == 2 {
			c.combinations[card] = 0
			pairsNumber++
		}
	}

	if threeOfAkind && pairsNumber == 1 {
		return FullHouse
	} else if threeOfAkind {
		return ThreeOfAKind
	} else if pairsNumber == 2 {
		return TwoPair
	} else if pairsNumber == 1 {
		return OnePair
	}

	return HighCard
}

func (c *cardFactory) parseCombination(line string) CardsCombination {
	var cards FiveCards

	parts := strings.Split(line, " ")
	assertTrue(len(parts) == 2, "non expected length")

	score, err := strconv.Atoi(parts[1])
	assertNoError(err)

	for i, cardSymbol := range []byte(parts[0]) {
		cardItem, ok := cardMap[cardSymbol]
		assertTrue(ok, "non expected card")
		cards[i] = cardItem
	}

	return CardsCombination{
		combination: c.getCombination(&cards),
		cards:       cards,
		score:       score,
	}
}

type CardsCombination struct {
	combination CardCombinationScore
	cards       FiveCards
	score       int
}

func (c *CardsCombination) Compare(other *CardsCombination) int {
	if c.combination > other.combination {
		return 1
	} else if c.combination < other.combination {
		return -1
	} else {
		return FiveCardsCompare(&c.cards, &other.cards)
	}
}

type FiveCards = [5]CardKind

func FiveCardsCompare(a *FiveCards, b *FiveCards) int {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		} else {
			continue
		}
	}

	return 0
}

func assertNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func assertTrue(cond bool, message string) {
	if !cond {
		panic(message)
	}
}
