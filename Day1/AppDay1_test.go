package Day1

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	data := Utils.ReadFileByLinesInt("\\masses.txt")
	expected := 3382284
	actual := Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 1 Part 1 Test 1")
	}

	data = []int{12}
	expected = 2
	actual = Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 1 Part 1 Test 2")
	}

	data = []int{14}
	expected = 2
	actual = Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 1 Part 1 Test 3")
	}

	data = []int{1969}
	expected = 654
	actual = Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 1 Part 1 Test 4")
	}

	data = []int{100756}
	expected = 33583
	actual = Day1Part1(data)
	if actual != expected {
		t.Errorf("Test Day 1 Part 1 Test 5")
	}
}

func TestDay1Part2(t *testing.T) {
	data := Utils.ReadFileByLinesInt("\\masses.txt")
	expected := 5070541
	actual := Day1Part2(data)
	if actual != expected {
		t.Errorf("Test Day 1 Part 2 Test 1")
	}

	data = []int{14}
	expected = 2
	actual = Day1Part2(data)
	if actual != expected {
		t.Errorf("Test Day 1 Part 2 Test 2")
	}

	data = []int{1969}
	expected = 966
	actual = Day1Part2(data)
	if actual != expected {
		t.Errorf("Test Day 1 Part 2 Test 3")
	}

	data = []int{100756}
	expected = 50346
	actual = Day1Part2(data)
	if actual != expected {
		t.Errorf("Test Day 1 Part 2 Test 4")
	}
}
