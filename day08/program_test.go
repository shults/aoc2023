package day08

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		{"test2.txt", 6},
		{"test.txt", 2},
	}

	for _, test := range tests {
		file, err := os.Open(test.file)
		panicOnError(err)

		program := newProgram(file)
		err = file.Close()
		panicOnError(err)

		res := program.Part1()

		assert.Equal(t, test.expected, res)
	}

}

func TestMyPart1(t *testing.T) {
	file, err := os.Open("my.txt")
	panicOnError(err)
	defer file.Close()

	program := newProgram(file)

	res := program.Part1()

	assert.Equal(t, 16343, res)
}

func TestExample2(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		{"test3.txt", 6},
	}

	for _, test := range tests {
		file, err := os.Open(test.file)
		panicOnError(err)

		program := newProgram(file)
		err = file.Close()
		panicOnError(err)

		res := program.Part2(true)

		assert.Equal(t, test.expected, res)
	}

}
