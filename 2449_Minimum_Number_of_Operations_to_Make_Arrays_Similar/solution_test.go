package leetcode

import (
	"sort"
	"testing"
)

func filter(nums []int) (evens []int, odds []int) {
	for i, num := range nums {
		if num%2 == 0 {
			evens = append(evens, i)
			continue
		}

		odds = append(odds, i)
	}

	return evens, odds
}

func solution(nums []int, target []int) int64 {
	sort.Ints(nums)
	sort.Ints(target)

	evens, odds := filter(nums)
	evensT, oddsT := filter(target)

	count := 0.0

	for i := 0; i < len(evens); i++ {
		difference := target[evensT[i]] - nums[evens[i]]

		if difference > 0 {
			count += float64(difference) / 2.0
		}
	}

	for i := 0; i < len(odds); i++ {
		difference := target[oddsT[i]] - nums[odds[i]]

		if difference > 0 {
			count += float64(difference) / 2.0
		}
	}

	return int64(count)
}

func TestSolution(t *testing.T) {
	result := solution([]int{8, 12, 6}, []int{2, 14, 10})

	if result != 2 {
		t.Errorf("wanted 2, got %d", result)
	}

	result2 := solution([]int{1, 2, 5}, []int{4, 1, 3})

	if result2 != 1 {
		t.Errorf("wanted 1, got %d", result2)
	}
}
