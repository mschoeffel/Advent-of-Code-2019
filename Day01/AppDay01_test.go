package Day01

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay01Part1Test1(t *testing.T) {
	data := Utils.ReadFileByLinesInt("\\masses.txt")
	expected := 3382284
	actual := Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 1 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay01Part1Test2(t *testing.T) {
	data := []int{12}
	expected := 2
	actual := Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 1 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}
func TestDay01Part1Test3(t *testing.T) {
	data := []int{14}
	expected := 2
	actual := Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 1 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay01Part1Test4(t *testing.T) {
	data := []int{1969}
	expected := 654
	actual := Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 1 Test 4: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay01Part1Test5(t *testing.T) {
	data := []int{100756}
	expected := 33583
	actual := Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 1 Test 5: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay01Part2Test1(t *testing.T) {
	data := Utils.ReadFileByLinesInt("\\masses.txt")
	expected := 5070541
	actual := Day1Part2(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 2 Test 1: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay01Part2Test2(t *testing.T) {
	data := []int{14}
	expected := 2
	actual := Day1Part2(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 2 Test 2: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay01Part2Test3(t *testing.T) {
	data := []int{1969}
	expected := 966
	actual := Day1Part2(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 2 Test 3: Expected: %v Actual: %v", expected, actual)
	}
}

func TestDay01Part2Test4(t *testing.T) {
	data := []int{100756}
	expected := 50346
	actual := Day1Part2(data)
	if actual != expected {
		t.Errorf("Test Day 01 Part 2 Test 4: Expected: %v Actual: %v", expected, actual)
	}
}
