package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardFactory(t *testing.T) {
	tests := []struct {
		in           string
		combination1 CardCombinationScore
		combination2 CardCombinationScore
	}{
		{"32T3K 765", OnePair, OnePair},
		{"T55J5 684", ThreeOfAKind, FourOfAKind},
		{"KK677 28", TwoPair, TwoPair},
		{"KTJJT 220", TwoPair, FourOfAKind},
		{"QQQJA 483", ThreeOfAKind, FourOfAKind},

		// my adds
		{"JJJJA 483", FourOfAKind, FiveOfAKind},
		{"JJJJJ 483", FiveOfAKind, FiveOfAKind},
		{"2345A 483", HighCard, HighCard},
		{"AAAAA 483", FiveOfAKind, FiveOfAKind},
		{"AAAAT 483", FourOfAKind, FourOfAKind},
		{"AAATT 483", FullHouse, FullHouse},
		{"AAA23 483", ThreeOfAKind, ThreeOfAKind},
		{"AA223 483", TwoPair, TwoPair},
		{"AA234 483", OnePair, OnePair},
		{"J22KA 483", OnePair, ThreeOfAKind},
	}

	cf := newCardFactory()

	for _, test := range tests {
		card := cf.parseCombination(test.in)
		assert.Equal(t, test.combination1, card.combination1, test.in)
		assert.Equal(t, test.combination2, card.combination2, test.in)
	}
}
