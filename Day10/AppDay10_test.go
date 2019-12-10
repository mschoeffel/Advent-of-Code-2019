package Day10

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay09Part1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\map.txt")
	expected := 299
	actual := Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 1")
	}

	data = Utils.ReadFileByLinesString("\\map_test1.txt")
	expected = 33
	actual = Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 2")
	}

	data = Utils.ReadFileByLinesString("\\map_test2.txt")
	expected = 35
	actual = Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 3")
	}

	data = Utils.ReadFileByLinesString("\\map_test3.txt")
	expected = 41
	actual = Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 4")
	}

	data = Utils.ReadFileByLinesString("\\map_test4.txt")
	expected = 210
	actual = Day10Part1(data)
	if actual != expected {
		t.Errorf("Test Day 10 Part 1 Test 5")
	}
}
