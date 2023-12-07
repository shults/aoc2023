package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardFactory(t *testing.T) {
	tests := []struct {
		in          string
		combination CardCombinationScore
	}{
		{"32T3K 765", OnePair},
		{"T55J5 684", ThreeOfAKind},
		{"KK677 28", TwoPair},
		{"KTJJT 220", TwoPair},
		{"QQQJA 483", ThreeOfAKind},

		// my adds
		{"2345A 483", HighCard},
		{"AAAAA 483", FiveOfAKind},
		{"AAAAT 483", FourOfAKind},
		{"AAATT 483", FullHouse},
		{"AAA23 483", ThreeOfAKind},
		{"AA223 483", TwoPair},
		{"AA234 483", OnePair},
	}

	cf := newCardFactory()

	for _, test := range tests {
		card := cf.parseCombination(test.in)
		assert.Equal(t, test.combination, card.combination)
	}
}
