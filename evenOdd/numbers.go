package main

import (
	"math"
	"strconv"
)

type numbers []int

func startNumbers() numbers {
	return numbers{}
}

func setEvenOrOdds(size int, evens bool) numbers {
	nums := numbers{}

	for i := 0; i <= size; i++ {
		if math.Mod(float64(i), 2) == 0 && evens {
			nums = append(nums, i)

		} else if math.Mod(float64(i), 2) != 0 && !evens {
			nums = append(nums, i)
		}

	}

	return nums
}

func (n numbers) evens(interval int) numbers {
	return setEvenOrOdds(interval, true)

}

func (n numbers) odds(interval int) numbers {
	return setEvenOrOdds(interval, false)

}

func (n numbers) toString() string {
	s := ""
	for _, toStr := range n {
		if s == "" {
			s = strconv.Itoa(toStr)
		} else {
			s = s + ", " + strconv.Itoa(toStr)

		}

	}
	return s
}
