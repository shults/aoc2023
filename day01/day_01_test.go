package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDigit(t *testing.T) {
	digits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	for _, digit := range digits {
		assert.True(t, isAsciiDigit(digit))
	}

	nonDigits := []byte("asdzxcooasdwemmc")

	for _, dn := range nonDigits {
		assert.False(t, isAsciiDigit(dn))
	}
}

func TestGetTwoDigitsNumber(t *testing.T) {
	// part 1
	assert.Equal(t, 12, GetTwoDigitsNumber("1abc2"))
	assert.Equal(t, 38, GetTwoDigitsNumber("pqr3stu8vwx"))
	assert.Equal(t, 15, GetTwoDigitsNumber("a1b2c3d4e5f"))
	assert.Equal(t, 77, GetTwoDigitsNumber("treb7uchet"))

	// part2
	assert.Equal(t, 29, GetTwoDigitsNumber("two1nine"))
	assert.Equal(t, 83, GetTwoDigitsNumber("eightwothree"))
	assert.Equal(t, 13, GetTwoDigitsNumber("abcone2threexyz"))
	assert.Equal(t, 24, GetTwoDigitsNumber("xtwone3four"))
	assert.Equal(t, 42, GetTwoDigitsNumber("4nineeightseven2"))
	assert.Equal(t, 14, GetTwoDigitsNumber("zoneight234"))
	assert.Equal(t, 76, GetTwoDigitsNumber("7pqrstsixteen"))
	assert.Equal(t, 22, GetTwoDigitsNumber("two"))
	assert.Equal(t, 23, GetTwoDigitsNumber("twothree"))
}
