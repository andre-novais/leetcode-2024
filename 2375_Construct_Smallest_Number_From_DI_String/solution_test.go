package leetcode

import (
	"testing"
)

func solution(pattern string) string {

	usedInts := []bool{false, false, false, false, false, false, false, false, false}

	return "1"
}

func getNextStreak(i int, patternRunes []rune) (char string, streak int) {
	charR := patternRunes[i]
	i++
	streak = 1

	for i < len(patternRunes) {
		next := patternRunes[i]
		if next != charR {
			return string(charR), streak
		}

		streak++
		i++
	}

	return string(charR), streak
}

func getNextInt(last int, indDec string, streak int, usedInts *[]bool) string {
	if incDec == "D" {
		for streak > 0 {
			for i := len(*usedInts) - 1 - streak; i > 0; i-- {

			}
		}
	}
}

func TestSolution(t *testing.T) {
	result := solution("IIIDIDDD")

	if result != "123849765" {
		t.Errorf("wanted 123849765, got %s", result)
	}

	result2 := solution("DDD")

	if result2 != "4321" {
		t.Errorf("wanted 2, got %s", result2)
	}
}
