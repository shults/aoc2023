package day02

import (
	"fmt"
	"os"
)

func Main() {
	fmt.Printf("Not implemented yet\n")
	os.Exit(1)
}

func ParseGame(gameLine string) Game {
	panic("impl")
}

type Game struct {
	id       int
	gameSets []GameSet
}

func (g *Game) Accepts(bag *Bag) bool {
	panic("impl")
}

type GameSet struct {
	red   int
	green int
	blue  int
}

func NewBag(red, green, blue int) Bag {
	return Bag{
		red:   red,
		green: green,
		blue:  blue,
	}
}

type Bag = GameSet
