package day09

import (
	"strconv"
	"strings"
)

func ParseSequence(numbers string) (*Sequence, error) {
	strNums := strings.Fields(numbers)
	nums := make([]int, len(strNums))

	for i, strNum := range strNums {
		num, err := strconv.Atoi(strNum)

		if err != nil {
			return nil, err
		}

		nums[i] = num
	}

	prevs := make([]int, 0)
	currSeq := make([]int, len(nums))
	copy(currSeq, nums)

	{
		i := 0
		for {
			i++
			allDiffsAreZero := true
			nextSeq := make([]int, len(currSeq)-1)

			for j := 0; j < len(currSeq)-1; j++ {
				nextSeq[j] = currSeq[j+1] - currSeq[j]

				if nextSeq[j] != 0 {
					allDiffsAreZero = false
				}
			}

			prevs = append(prevs, currSeq[len(currSeq)-2])

			if allDiffsAreZero {
				break
			}

			currSeq = nextSeq
		}
	}

	return &Sequence{
		prev:    prevs,
		current: nums[len(nums)-1],
	}, nil
}

func NewSequence(current int, prev []int) Sequence {
	return Sequence{
		prev:    prev,
		current: current,
	}
}

type Sequence struct {
	prev    []int
	current int
}

func (s *Sequence) GetNext() int {
	return s.GetByShift(1)
}

func (s *Sequence) GetByShift(shift uint) int {
	if shift == 0 {
		return s.current
	}

	prev := make([]int, len(s.prev))
	copy(prev, s.prev)

	current := s.current
	for i := 0; i < int(shift); i++ {
		current = s.genNext(current, prev, 0)
	}

	return current
}

func (s *Sequence) genNext(currentVal int, previous []int, depth int) int {
	if len(previous)-1 == depth {
		return previous[depth]
	}

	res := currentVal + s.genNext(currentVal-previous[depth], previous, depth+1)
	previous[depth] = currentVal

	return res
}
