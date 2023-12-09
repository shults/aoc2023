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

			endList = append(endList, currSeq[len(currSeq)-2])
			startList = append(startList, currSeq[0])

			if allDiffsAreZero {
				break
			}

			currSeq = nextSeq
		}
	}

	return &Sequence{
		end:       nums[len(nums)-1],
		startList: startList,
		endList:   endList,
	}, nil
}

func NewSequence(
	end int,
	startList, endList []int,
) Sequence {
	return Sequence{
		end:       end,
		startList: startList,
		endList:   endList,
	}
}

type Sequence struct {
	end                int
	startList, endList []int
}

func (s *Sequence) GetNext() int {
	return s.GetNextByShift(1)
}

func (s *Sequence) GetPrev() int {
	return s.GetPrevByShift(1)
}

func (s *Sequence) GetPrevByShift(shift uint) int {
	if shift == 0 {
		return s.startList[0]
	}

	startList := make([]int, len(s.startList))
	copy(startList, s.startList)

	start := 0
	for i := 0; i < int(shift); i++ {
		start = s.genPrev(startList)
	}

	return start
}

func (s *Sequence) genPrev(startList []int) int {
	last := startList[len(startList)-1]

	for i := len(startList) - 2; i > -1; i-- {
		last = startList[i] - last
		startList[i] = last
	}

	return last
}

func (s *Sequence) GetNextByShift(shift uint) int {
	if shift == 0 {
		return s.end
	}

	endList := make([]int, len(s.endList))
	copy(endList, s.endList)

	last := s.end
	for i := 0; i < int(shift); i++ {
		last = s.genNext(last, endList, 0)
	}

	return last
}

func (s *Sequence) genNext(last int, endList []int, depth int) int {
	if len(endList)-1 == depth {
		return endList[depth]
	}

	res := last + s.genNext(last-endList[depth], endList, depth+1)
	endList[depth] = last

	return res
}
