package main

import (
	"Advent-of-Code-2019/Day4"
)

func testDay4Part1() {
	min := 246515
	max := 739105
	expected := 1048
	actual := Day4.Day4Part1(min, max)
	if actual == expected {
		writeSuccessMesage(4, 1, 1)
	} else {
		writeErrorMessageInt(4, 1, 1, expected, actual)
	}
}

func testDay4Part2() {
	min := 246515
	max := 739105
	expected := 677
	actual := Day4.Day4Part2(min, max)
	if actual == expected {
		writeSuccessMesage(4, 2, 1)
	} else {
		writeErrorMessageInt(4, 2, 1, expected, actual)
	}
}
