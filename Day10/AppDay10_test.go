package Day10

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay10Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map.txt")
	expected := 299
	actual := Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay10Part1Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map_test1.txt")
	expected := 33
	actual := Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay10Part1Test3(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map_test2.txt")
	expected := 35
	actual := Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay10Part1Test4(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map_test3.txt")
	expected := 41
	actual := Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 4: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay10Part1Test5(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map_test4.txt")
	expected := 210
	actual := Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 5: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay10Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map.txt")
	expected := 1419
	actual := Day10Part2(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay10Part2Test2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map_test4.txt")
	expected := 802
	actual := Day10Part2(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 2 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}
