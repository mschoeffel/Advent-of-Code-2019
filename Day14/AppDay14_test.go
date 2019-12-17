package Day14

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay14Part1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\processes.txt")
	expected := 907302
	actual := Day14Part1(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 1 Test 1")
	}

	data = Utils.ReadFileByLinesString("\\processes_test1.txt")
	expected = 165
	actual = Day14Part1(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 1 Test 2")
	}

	data = Utils.ReadFileByLinesString("\\processes_test2.txt")
	expected = 13312
	actual = Day14Part1(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 1 Test 3")
	}

	data = Utils.ReadFileByLinesString("\\processes_test3.txt")
	expected = 180697
	actual = Day14Part1(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 1 Test 4")
	}

	data = Utils.ReadFileByLinesString("\\processes_test4.txt")
	expected = 2210736
	actual = Day14Part1(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 1 Test 5")
	}
}

func TestDay14Part2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\processes.txt")
	expected := 1670299
	actual := Day14Part2(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 2 Test 1")
	}

	data = Utils.ReadFileByLinesString("\\processes_test1.txt")
	expected = 6323777403
	actual = Day14Part2(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 2 Test 2")
	}

	data = Utils.ReadFileByLinesString("\\processes_test2.txt")
	expected = 82892753
	actual = Day14Part2(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 2 Test 3")
	}

	data = Utils.ReadFileByLinesString("\\processes_test3.txt")
	expected = 5586022
	actual = Day14Part2(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 2 Test 4")
	}

	data = Utils.ReadFileByLinesString("\\processes_test4.txt")
	expected = 460664
	actual = Day14Part2(data)
	if actual != expected {
		t.Errorf("Test Day 14 Part 2 Test 5")
	}
}
