package main

import (
	"Advent-of-Code-2019/Day1"
	"Advent-of-Code-2019/Utils"
)

func testDay1Part1() {
	masses := Utils.ReadFileByLinesInt("\\Day1\\masses.txt")
	expected := 3382284
	actual := Day1.Day1Part1(masses)
	if actual == expected {
		writeSuccessMesage(1, 1, 1)
	} else {
		writeErrorMessageInt(1, 1, 1, expected, actual)
	}

	expected = 2
	actual = Day1.Day1Part1([]int{12})
	if actual == expected {
		writeSuccessMesage(1, 1, 2)
	} else {
		writeErrorMessageInt(1, 1, 2, expected, actual)
	}

	expected = 2
	actual = Day1.Day1Part1([]int{14})
	if actual == expected {
		writeSuccessMesage(1, 1, 3)
	} else {
		writeErrorMessageInt(1, 1, 3, expected, actual)
	}

	expected = 654
	actual = Day1.Day1Part1([]int{1969})
	if actual == expected {
		writeSuccessMesage(1, 1, 4)
	} else {
		writeErrorMessageInt(1, 1, 4, expected, actual)
	}

	expected = 33583
	actual = Day1.Day1Part1([]int{100756})
	if actual == expected {
		writeSuccessMesage(1, 1, 5)
	} else {
		writeErrorMessageInt(1, 1, 5, expected, actual)
	}
}

func testDay1Part2() {
	masses := Utils.ReadFileByLinesInt("\\Day1\\masses.txt")
	expected := 5070541
	actual := Day1.Day1Part2(masses)
	if actual == expected {
		writeSuccessMesage(1, 2, 1)
	} else {
		writeErrorMessageInt(1, 2, 1, expected, actual)
	}

	expected = 2
	actual = Day1.Day1Part2([]int{14})
	if actual == expected {
		writeSuccessMesage(1, 2, 2)
	} else {
		writeErrorMessageInt(1, 2, 2, expected, actual)
	}

	expected = 966
	actual = Day1.Day1Part2([]int{1969})
	if actual == expected {
		writeSuccessMesage(1, 2, 3)
	} else {
		writeErrorMessageInt(1, 2, 3, expected, actual)
	}

	expected = 50346
	actual = Day1.Day1Part2([]int{100756})
	if actual == expected {
		writeSuccessMesage(1, 2, 4)
	} else {
		writeErrorMessageInt(1, 2, 4, expected, actual)
	}
}
