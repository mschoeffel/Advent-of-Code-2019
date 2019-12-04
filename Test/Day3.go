package main

import (
	"Advent-of-Code-2019/Day3"
	"Advent-of-Code-2019/Utils"
)

func testDay3Part1() {
	data := Utils.ReadFileByLinesString("\\Day3\\wires.txt")
	expected := 248
	actual := Day3.Day3Part1(data)
	if actual == expected {
		writeSuccessMesage(3, 1, 1)
	} else {
		writeErrorMessageInt(3, 1, 1, expected, actual)
	}

	data = []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}
	expected = 6
	actual = Day3.Day3Part1(data)
	if actual == expected {
		writeSuccessMesage(3, 1, 2)
	} else {
		writeErrorMessageInt(3, 1, 2, expected, actual)
	}

	data = []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	expected = 159
	actual = Day3.Day3Part1(data)
	if actual == expected {
		writeSuccessMesage(3, 1, 3)
	} else {
		writeErrorMessageInt(3, 1, 3, expected, actual)
	}

	data = []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	expected = 135
	actual = Day3.Day3Part1(data)
	if actual == expected {
		writeSuccessMesage(3, 1, 4)
	} else {
		writeErrorMessageInt(3, 1, 4, expected, actual)
	}
}

func testDay3Part2() {
	data := Utils.ReadFileByLinesString("\\Day3\\wires.txt")
	expected := 28580
	actual := Day3.Day3Part2(data)
	if actual == expected {
		writeSuccessMesage(3, 2, 1)
	} else {
		writeErrorMessageInt(3, 2, 1, expected, actual)
	}

	data = []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}
	expected = 30
	actual = Day3.Day3Part2(data)
	if actual == expected {
		writeSuccessMesage(3, 2, 2)
	} else {
		writeErrorMessageInt(3, 2, 2, expected, actual)
	}

	data = []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	expected = 610
	actual = Day3.Day3Part2(data)
	if actual == expected {
		writeSuccessMesage(3, 2, 3)
	} else {
		writeErrorMessageInt(3, 2, 3, expected, actual)
	}

	data = []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}
	expected = 410
	actual = Day3.Day3Part2(data)
	if actual == expected {
		writeSuccessMesage(3, 2, 4)
	} else {
		writeErrorMessageInt(3, 2, 4, expected, actual)
	}
}
