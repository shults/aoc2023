package main

import (
	"aoc2023/day01"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var res int = 0

	for {
		line, _, err := r.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("err=%s\n", err)
			os.Exit(1)
		}

		res += day01.GetTwoDigitsNumber(string(line))
	}

	fmt.Printf("%d\n", res)
}
