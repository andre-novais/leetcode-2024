package leetcode

import (
	"sort"
	"testing"
)

func solution(childrenGreed []int, cookies []int) int {
	sort.Ints(childrenGreed)
	sort.Ints(cookies)

	g := len(childrenGreed) - 1
	c := len(cookies) - 1
	count := 0

	for g > -1 && c > -1 {
		if childrenGreed[g] <= cookies[c] {
			count++
			c--
			g--
			continue
		}

		g--
	}

	return count
}

func TestSolution(t *testing.T) {
	result := solution([]int{1, 2, 3}, []int{1, 1})

	if result != 1 {
		t.Errorf("wanted 1, got %d", result)
	}

	result2 := solution([]int{1, 2}, []int{1, 2, 3})

	if result2 != 2 {
		t.Errorf("wanted 2, got %d", result2)
	}
}
