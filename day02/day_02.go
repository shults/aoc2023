package day02

import (
	"aoc2023/tools"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	colorRed   = "red"
	colorBlue  = "blue"
	colorGreen = "green"
)

func Main(args []string, in io.Reader) {
	r := bufio.NewReader(in)
	var red, green, blue int
	var help bool

	flagSet := flag.NewFlagSet("day2", flag.ExitOnError)
	flagSet.IntVar(&red, "r", 0, "number of red")
	flagSet.IntVar(&green, "g", 0, "number of green")
	flagSet.IntVar(&blue, "b", 0, "number of blue")
	flagSet.BoolVar(&help, "help", false, "print help")

	err := flagSet.Parse(args)

	if err != nil {
		flagSet.Usage()
		fmt.Printf("Got error: %s\n", err)
		os.Exit(1)
	}

	if help {
		flagSet.Usage()
		os.Exit(0)
	}

	bag, err := NewBagFromArgs(red, green, blue)

	if err != nil {
		flagSet.Usage()
		fmt.Printf("Got error: %s\n", err)
		os.Exit(1)
	}

	var acceptableGames = 0
	var sumOfPower = 0

	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("err=%s\n", err)
			os.Exit(1)
		}

		game, err := ParseGame(string(line))

		if err != nil {
			fmt.Printf("faile to parse game: %s\n", err)
			os.Exit(1)
		}

		sumOfPower += game.Power()

		if game.Accepts(&bag) {
			acceptableGames += game.id
		}
	}

	fmt.Printf("acceptableGames = %d sumOfPower = %d\n", acceptableGames, sumOfPower)
}

func ParseGame(gameLine string) (*Game, error) {
	var game Game

	split := strings.Split(gameLine, ":")

	if len(split) != 2 {
		return nil, fmt.Errorf("wrong format: non expected number of parts in split %+v", split)
	}

	gameData := strings.Split(split[0], " ")

	if len(gameData) != 2 {
		return nil, fmt.Errorf("wrong format: unexpected game format %+v", gameData)
	}

	id, err := strconv.Atoi(gameData[1])

	if err != nil {
		return nil, err
	}

	game.id = id
	gameSets, err := parseGameSets(split[1])

	if err != nil {
		return nil, err
	}

	game.AppendGameSets(gameSets...)

	return &game, nil
}

func parseGameSets(gameSetStr string) ([]GameSet, error) {
	return tools.Map(strings.Split(gameSetStr, ";"), func(gameSetNotTrimmed *string) (GameSet, error) {
		gameSet := NewGameSet(0, 0, 0)

		err := tools.Each(strings.Split(strings.TrimSpace(*gameSetNotTrimmed), ","), func(colorPairStr *string) error {
			pair := strings.Split(strings.TrimSpace(*colorPairStr), " ")

			if len(pair) != 2 {
				return fmt.Errorf("unable to parse color pair: %+v", pair)
			}

			num, err := strconv.Atoi(pair[0])

			if err != nil {
				return fmt.Errorf("unable to parse color pair: %s", err)
			}

			if num < 0 {
				return fmt.Errorf("unable to parse color pair: negative value %d", num)
			}

			switch pair[1] {
			case colorRed:
				gameSet.red += num
			case colorGreen:
				gameSet.green += num
			case colorBlue:
				gameSet.blue += num
			default:
				return fmt.Errorf("unable to parse color pair: negative unknown color %s", pair[1])
			}

			return nil
		})

		if err != nil {
			return gameSet, err
		}

		return gameSet, nil
	})
}

type Game struct {
	id       int
	gameSets []GameSet
	power    GameSet
}

func (g *Game) AppendGameSets(sets ...GameSet) {
	for _, gs := range sets {
		g.appendOne(gs)
	}
}

func (g *Game) appendOne(set GameSet) {
	g.gameSets = append(g.gameSets, set)

	g.power.red = max(g.power.red, set.red)
	g.power.green = max(g.power.green, set.green)
	g.power.blue = max(g.power.blue, set.blue)
}

func (g *Game) Accepts(bag *Bag) bool {
	return tools.MustEvery(g.gameSets, func(set *GameSet) bool {
		return set.Accepts(bag)
	})
}

func (g *Game) Power() int {
	return g.power.red * g.power.green * g.power.blue
}

func NewGameSet(red, green, blue int) GameSet {
	return NewBag(red, green, blue)
}

type GameSet struct {
	red   int
	green int
	blue  int
}

func (gs *GameSet) Accepts(bag *Bag) bool {
	return gs.red <= bag.red && gs.green <= bag.green && gs.blue <= bag.blue
}

func NewBag(red, green, blue int) Bag {
	return Bag{
		red:   red,
		green: green,
		blue:  blue,
	}
}

func NewBagFromArgs(red, green, blue int) (Bag, error) {
	bag := NewBag(red, green, blue)

	if bag.red < 0 || bag.green < 0 || bag.blue < 0 {
		return bag, fmt.Errorf("Unable to parse bag from cli. Bag cannot contain negative number of items %+v\n", bag)
	}

	return bag, nil
}

type Bag = GameSet

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
