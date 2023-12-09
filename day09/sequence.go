package day09

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseSequence(numbers string) (*Sequence, error) {
	strNums := strings.Fields(numbers)
	nums := make([]int, len(strNums))

	for i, strNum := range strNums {
		num, err := strconv.Atoi(strNum)

		if err != nil {
			return nil, fmt.Errorf("failed to parse int: %s", err)
		}

		nums[i] = num
	}

	endList := make([]int, 0)
	startList := make([]int, 0)
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

			endList = append(endList, currSeq[len(currSeq)-1])
			startList = append(startList, currSeq[0])

			if allDiffsAreZero {
				break
			}

			currSeq = nextSeq
		}
	}

	return &Sequence{
		startList: startList,
		endList:   endList,
	}, nil
}

func NewSequence(
	startList, endList []int,
) Sequence {
	return Sequence{
		startList: startList,
		endList:   endList,
	}
}

type Sequence struct {
	startList, endList []int
}

func (s *Sequence) GetNext() int {
	return s.GetNextByShift(1)
}

func (s *Sequence) GetPrev() int {
	return s.GetPrevByShift(1)
}

func (s *Sequence) GetPrevByShift(shift uint) int {
	return s.getDayByShift(shift, s.startList, true)
}

func (s *Sequence) GetNextByShift(shift uint) int {
	return s.getDayByShift(shift, s.endList, false)
}

func (s *Sequence) getDayByShift(shift uint, data []int, start bool) int {
	if shift == 0 {
		return data[0]
	}

	dataCopy := make([]int, len(data))
	copy(dataCopy, data)

	for i := 0; i < int(shift); i++ {
		for j := len(dataCopy) - 2; j > -1; j-- {
			if start {
				dataCopy[j] -= dataCopy[j+1]
			} else {
				dataCopy[j] += dataCopy[j+1]
			}
		}
	}

	return dataCopy[0]
}
