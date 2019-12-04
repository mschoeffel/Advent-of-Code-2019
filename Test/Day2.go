package main

import (
	"Advent-of-Code-2019/Day2"
	"Advent-of-Code-2019/Utils"
)

func testDay2Part1() {
	data := Utils.ReadFileBySeparatorInt("\\Day2\\codes.txt", ",")
	expected := 3267740
	actual := Day2.Day2Part1(data, 12, 2)
	if actual == expected {
		writeSuccessMesage(2, 1, 1)
	} else {
		writeErrorMessageInt(2, 1, 1, expected, actual)
	}
}

func testDay2Part2() {
	data := 19690720
	expected := 7870
	actual := Day2.Day2Part2(data)
	if actual == expected {
		writeSuccessMesage(2, 2, 1)
	} else {
		writeErrorMessageInt(2, 2, 1, expected, actual)
	}
}
