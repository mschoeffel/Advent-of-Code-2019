package Day09

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay09Part1Test1(t *testing.T) {
	data := Utils.ReadFileBySeparatorInt("\\codes.txt", ",")
	expected := 2594708277
	actual := Day9Part1(data, "1", true)
	if actual != expected {
		t.Errorf("Test Day 09 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay09Part1Test2(t *testing.T) {
	data := []int{104, 1125899906842624, 99}
	expected := 1125899906842624
	actual := Day9Part1(data, "0", true)
	if actual != expected {
		t.Errorf("Test Day 09 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay09Part1Test3(t *testing.T) {
	data := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	expected := 1219070632396864
	actual := Day9Part1(data, "0", true)
	if actual != expected {
		t.Errorf("Test Day 09 Part 1 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay09Part1Test4(t *testing.T) {
	data := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	expected := 99
	actual := Day9Part1(data, "0", true)
	if actual != expected {
		t.Errorf("Test Day 09 Part 1 Test 4: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay09Part2Test1(t *testing.T) {
	data := Utils.ReadFileBySeparatorInt("\\codes.txt", ",")
	expected := 87721
	actual := Day9Part1(data, "2", true)
	if actual != expected {
		t.Errorf("Test Day 09 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}
