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

func ReadFile(name string) ([]string, error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	var res []string
	line := ""
	for {
		bytes, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		line += string(bytes)

		if isPrefix {
			continue
		}

		res = append(res, line)
		line = ""
	}

	return res, nil
}
