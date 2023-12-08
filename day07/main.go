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
	CardJ
	CardT
	Card9
	Card8
	Card7
	Card6
	Card5
	Card4
	Card3
	Card2
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

	invertedCardMap map[CardKind]byte
)

func init() {
	invertedCardMap = make(map[CardKind]byte, len(cardMap))

	for k, v := range cardMap {
		invertedCardMap[v] = k
	}
}

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
		file, err := os.Open(*inputFile)

		if err != nil {
			flagSet.Usage()
			fmt.Printf("Error ocurred: %s\n", err)
			os.Exit(1)
		}

		defer file.Close()
		in = file
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
	cc := newCardCombinations()

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		assertNoError(err)
		cc.addCard(string(line))
	}

	calcTimeStart = time.Now()

	part1 = cc.part1()
	part2 = cc.part2()

	return
}

func newCardCombinations() *cardCombinations {
	return &cardCombinations{}
}

type cardCombinations struct {
	cf           cardFactory
	combinations []CardsCombination
}

func (cc *cardCombinations) addCard(line string) {
	cc.combinations = append(cc.combinations, cc.cf.parseCombination(line))
}

func (cc *cardCombinations) part1() int {
	part1 := 0

	slices.SortFunc(cc.combinations, func(a, b CardsCombination) int {
		return a.ComparePart1(&b)
	})
	for i, c := range cc.combinations {
		part1 += (i + 1) * c.score
	}

	return part1
}

func (cc *cardCombinations) part2() int {
	part2 := 0

	slices.SortFunc(cc.combinations, func(a, b CardsCombination) int {
		return a.ComparePart2(&b)
	})
	for i, c := range cc.combinations {
		fmt.Printf("%s\n", c.Debug2(i+1))
		part2 += (i + 1) * c.score
	}

	return part2
}

func newCardFactory() cardFactory {
	return cardFactory{}
}

type cardFactory struct {
	combinations [13]uint8
}

func (c *cardFactory) recalculateCombinations(cards *FiveCards) {
	for i := 0; i < len(c.combinations); i++ {
		c.combinations[i] = 0
	}
	for _, card := range cards {
		c.combinations[card] += 1
	}
}

func (c *cardFactory) getCombinationPart1(cards *FiveCards) CardCombinationScore {
	c.recalculateCombinations(cards)

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

func maxCombination(a, b CardCombinationScore) CardCombinationScore {
	if a > b {
		return a
	} else {
		return b
	}
}

func (c *cardFactory) getCombinationPart2(cards *FiveCards) CardCombinationScore {
	c.recalculateCombinations(cards)

	threeOfAkind := false
	pairsNumber := 0
	maxCombo := HighCard

	for _, card := range cards {
		if card == CardJ {
			continue
		}

		if c.combinations[card] == 5 {
			return FiveOfAKind
		} else if c.combinations[card] == 4 {
			maxCombo = maxCombination(maxCombo, FourOfAKind)
		} else if c.combinations[card] == 3 {
			maxCombo = maxCombination(maxCombo, ThreeOfAKind)
			c.combinations[card] = 0
			threeOfAkind = true
		} else if c.combinations[card] == 2 {
			maxCombo = maxCombination(maxCombo, OnePair)
			c.combinations[card] = 0
			pairsNumber++
		}
	}

	if pairsNumber == 2 {
		maxCombo = TwoPair
	}

	if threeOfAkind && pairsNumber == 1 {
		maxCombo = FullHouse
	}

	jokers := c.combinations[CardJ]
	if jokers > 0 {

		if jokers == 4 || jokers == 5 {
			// 5 | 4 -> 5 of kind | 0 hops
			return FiveOfAKind
		} else if jokers == 3 && pairsNumber > 0 {
			// 3 -> five of kind if pair | 5 hops pair -> two pairs -> three of kine -> full house -> 4k -> 5k 3
			return FiveOfAKind
		} else if jokers == 3 {
			// 3 and high card
			return FourOfAKind
		} else if jokers == 2 && maxCombo == ThreeOfAKind {
			return FiveOfAKind
		} else if jokers == 2 && maxCombo == OnePair {
			return FourOfAKind
		} else if jokers == 2 && maxCombo == HighCard {
			return ThreeOfAKind
		} else if jokers == 1 {
			if maxCombo == HighCard {
				return OnePair
			} else if maxCombo == OnePair {
				return ThreeOfAKind
			} else if maxCombo == TwoPair {
				return FullHouse
			} else if maxCombo == ThreeOfAKind {
				return FourOfAKind
			} else if maxCombo == FourOfAKind {
				return FiveOfAKind
			}

			panic(fmt.Sprintf("non expected combo %+v", c.combinations))
		} else {
			panic(fmt.Sprintf("non expected combo 2 %+v", c.combinations))
		}
	}

	return maxCombo
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
		combination1: c.getCombinationPart1(&cards),
		combination2: c.getCombinationPart2(&cards),
		cards:        cards,
		score:        score,
	}
}

type CardsCombination struct {
	combination1 CardCombinationScore
	combination2 CardCombinationScore
	cards        FiveCards
	score        int
}

var printCardBuf = make([]byte, 5)

func PrintCards(fc *FiveCards) string {
	for i, card := range fc {
		printCardBuf[i] = invertedCardMap[card]
	}

	return string(printCardBuf)
}

func PrintCardCombinationScore(c CardCombinationScore) string {
	switch c {

	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case FiveOfAKind:
		return "FiveOfAKind"

	default:
		panic("non expected card")
	}
}

func (c *CardsCombination) Debug2(rank int) string {
	return fmt.Sprintf("%d %s %s", rank, PrintCardCombinationScore(c.combination2), PrintCards(&c.cards))
}

func (c *CardsCombination) ComparePart1(other *CardsCombination) int {
	if c.combination1 > other.combination1 {
		return 1
	} else if c.combination1 < other.combination1 {
		return -1
	} else {
		return FiveCardsCompareP1(&c.cards, &other.cards)
	}
}

func (c *CardsCombination) ComparePart2(other *CardsCombination) int {
	if c.combination2 > other.combination2 {
		return 1
	} else if c.combination2 < other.combination2 {
		return -1
	} else {
		return FiveCardsCompareP2(&c.cards, &other.cards)
	}
}

type FiveCards = [5]CardKind

func FiveCardsCompareP1(a *FiveCards, b *FiveCards) int {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}

	return 0
}

func FiveCardsCompareP2(a *FiveCards, b *FiveCards) int {
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}

		if a[i] == CardJ {
			return -1
		}

		if b[i] == CardJ {
			return 1
		}

		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
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
