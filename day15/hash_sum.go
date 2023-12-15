package day15

import (
	"bufio"
	"io"
)

func HashSumPart1(in io.Reader) int {
	reader := bufio.NewReader(in)
	sum := 0

	currentValue := 0
	for {
		b, err := reader.ReadByte()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		if b == ',' {
			sum += currentValue
			currentValue = 0
			continue
		}

		currentValue += int(b)
		currentValue *= 17
		currentValue %= 256
	}

	sum += currentValue

	return sum
}
