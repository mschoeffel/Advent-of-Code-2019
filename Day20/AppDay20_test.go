package Day20

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay20Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze.txt")
	expected := 588
	actual := Day20Part1(data)
	if actual != expected {
		t.Errorf("Test Day 20 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay20Part1Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze_test1.txt")
	expected := 23
	actual := Day20Part1(data)
	if actual != expected {
		t.Errorf("Test Day 20 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay20Part1Test3(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\maze_test2.txt")
	expected := 58
	actual := Day20Part1(data)
	if actual != expected {
		t.Errorf("Test Day 20 Part 1 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}
