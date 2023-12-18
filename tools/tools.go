package tools

import (
	"bufio"
	"cmp"
	"io"
	"os"
)

func AssertTrue(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func AssertNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func ReadLinesFileLines(name string) (lines []string, err error) {
	file, err := os.Open(name)

	if err != nil {
		return
	}

	defer file.Close()

	return ReadLines(file)
}

func ReadLines(in io.Reader) (lines []string, err error) {
	reader := bufio.NewReader(in)

	line := ""
	for {
		bytes, isPrefix, e := reader.ReadLine()

		if e == io.EOF {
			break
		}

		if e != nil {
			return nil, e
		}

		line += string(bytes)

		if isPrefix {
			continue
		}

		lines = append(lines, line)
		line = ""
	}

	return
}

func IsAsciiNumber(symbol byte) bool {
	return symbol >= '0' && symbol <= '9'
}

func AsciiNumberToInt(symbol byte) int {
	return int(symbol - '0')
}
