package Day04

import (
	"fmt"
	"strconv"
)

func Main() {
	min := 246515
	max := 739105
	fmt.Println("Day 4: Result part one: " + strconv.Itoa(Day4Part1(min, max)))
	fmt.Println("Day 4: Result part two: " + strconv.Itoa(Day4Part2(min, max)))
}

func Day4Part1(min int, max int) int {
	valid := 0
	for i := min; i <= max; i++ {
		passwordRunes := []rune(strconv.Itoa(i))

		doubleContained := false
		isValid := true
		for j := 1; j < len(passwordRunes); j++ {
			if passwordRunes[j] < passwordRunes[j-1] {
				isValid = false
				break
			}
			if passwordRunes[j] == passwordRunes[j-1] {
				doubleContained = true
			}
		}
		if doubleContained && isValid {
			valid++
		}
	}
	return valid
}

func Day4Part2(min int, max int) int {
	valid := 0
	for i := min; i <= max; i++ {
		passwordRunes := []rune(strconv.Itoa(i))

		doubleContained := false
		isValid := true
		for j := 1; j < len(passwordRunes); j++ {
			if passwordRunes[j] < passwordRunes[j-1] {
				isValid = false
				break
			}
			if passwordRunes[j] == passwordRunes[j-1] {
				if j == len(passwordRunes)-1 || passwordRunes[j+1] != passwordRunes[j] {
					if j == 1 || passwordRunes[j-2] != passwordRunes[j-1] {
						doubleContained = true
					}
				}
			}
		}
		if doubleContained && isValid {
			valid++
		}
	}
	return valid
}
