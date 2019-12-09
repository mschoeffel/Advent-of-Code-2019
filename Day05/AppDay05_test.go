package Day05

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay05Part1(t *testing.T) {
	data := Utils.ReadFileBySeparatorInt("\\codes.txt", ",")
	expected := 4887191
	actual := Day5Part1(data, "1")
	if actual != expected {
		t.Errorf("Test Day 05 Part 1 Test 1")
	}
}

func TestDay05Part2(t *testing.T) {
	data := Utils.ReadFileBySeparatorInt("\\codes.txt", ",")
	expected := 3419022
	actual := Day5Part2(data, "5")
	if actual != expected {
		t.Errorf("Test Day 05 Part 2 Test 1")
	}

	data = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	expected = 999
	actual = Day5Part2(data, "7")
	if actual != expected {
		t.Errorf("Test Day 05 Part 2 Test 2")
	}

	data = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	expected = 999
	actual = Day5Part2(data, "4")
	if actual != expected {
		t.Errorf("Test Day 05 Part 2 Test 3")
	}

	data = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	expected = 1000
	actual = Day5Part2(data, "8")
	if actual != expected {
		t.Errorf("Test Day 05 Part 2 Test 4")
	}

	data = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	expected = 1001
	actual = Day5Part2(data, "9")
	if actual != expected {
		t.Errorf("Test Day 05 Part 2 Test 5")
	}

	data = []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	expected = 1001
	actual = Day5Part2(data, "11")
	if actual != expected {
		t.Errorf("Test Day 05 Part 2 Test 6")
	}

}
