package day14

func NewCycleDetector() CycleDetector {
	return CycleDetector{
		requiredMatches: 3,
	}
}

type CycleDetector struct {
	items           []int
	requiredMatches int
	cyclePrefix     []int
	cycle           []int
}

func (c *CycleDetector) AddAndTryDetectCycle(items ...int) bool {
	if len(items) > 1 {
		for _, item := range items {
			if c.AddAndTryDetectCycle(item) {
				return true
			}
		}
		return false
	}

	c.items = append(c.items, items[0])

	return c.detectCycle()
}

func (c *CycleDetector) detectCycle() bool {
	if c.cycle != nil {
		return true
	}

	if len(c.items) < c.requiredMatches {
		return false
	}

	for i := 1; i < len(c.items)/c.requiredMatches; i++ {
		if c.detectCycleByLength(i) {
			return true
		}
	}

	return false
}

func (c *CycleDetector) Predict(index int) int {
	if index < 0 {
		panic("unable to make backward prediction")
	}

	if index <= len(c.cyclePrefix)-1 {
		return c.cyclePrefix[index]
	}

	index -= len(c.cyclePrefix)

	index = index % len(c.cycle)

	return c.cycle[index]
}

func (c *CycleDetector) detectCycleByLength(length int) bool {
	maybeCycle := c.items[len(c.items)-length:]

	for i := 1; i < c.requiredMatches; i++ {
		otherCycle := c.items[len(c.items)-length*(i+1) : len(c.items)-length*i]

		for j := 0; j < len(otherCycle); j++ {
			if otherCycle[j] != maybeCycle[j] {
				return false
			}
		}
	}

	c.cycle = maybeCycle

loop:
	for i := 0; ; i++ {
		other := c.items[i : len(maybeCycle)+i]

		for j := 0; j < len(other); j++ {
			if other[j] != maybeCycle[j] {
				continue loop
			}
		}

		c.cyclePrefix = c.items[:i]
		break
	}

	return true
}
