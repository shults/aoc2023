package day15

func Hash(str string) int {
	currentValue := 0

	for _, sym := range []byte(str) {
		currentValue += int(sym)
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}
