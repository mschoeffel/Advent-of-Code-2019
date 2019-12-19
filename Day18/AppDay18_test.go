package Day18

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay18Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze.txt")
	expected := 3918
	actual := Day18Part1(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part1Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze_test1.txt")
	expected := 8
	actual := Day18Part1(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part1Test3(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze_test2.txt")
	expected := 86
	actual := Day18Part1(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 1 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part1Test4(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze_test3.txt")
	expected := 132
	actual := Day18Part1(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 1 Test 4: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part1Test5(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze_test4.txt")
	expected := 136
	actual := Day18Part1(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 1 Test 5: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part1Test6(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze_test5.txt")
	expected := 81
	actual := Day18Part1(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 1 Test 6: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze2.txt")
	expected := 2004
	actual := Day18Part2(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part2Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze2_test1.txt")
	expected := 8
	actual := Day18Part2(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 2 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part2Test3(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze2_test2.txt")
	expected := 24
	actual := Day18Part2(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 2 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part2Test4(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze2_test3.txt")
	expected := 32
	actual := Day18Part2(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 2 Test 4: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay18Part2Test5(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze2_test4.txt")
	expected := 72
	actual := Day18Part2(data)
	if actual != expected {
		t.Errorf("Test Day 18 Part 2 Test 5: Expected: %v Actual: %v", expected, actual)
	}
}
