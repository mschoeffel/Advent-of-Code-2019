package Day16

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay16Part1(t *testing.T) {
	data := Utils.ReadFileByCharInt("\\inputs.txt")
	expected := "34841690"
	actual := Day16Part1(data)
	if actual != expected {
		t.Errorf("Test Day 16 Part 1 Test 1")
	}

	data = Utils.ReadFileByCharInt("\\inputs_test1.txt")
	expected = "24176176"
	actual = Day16Part1(data)
	if actual != expected {
		t.Errorf("Test Day 16 Part 1 Test 2")
	}

	data = Utils.ReadFileByCharInt("\\inputs_test2.txt")
	expected = "73745418"
	actual = Day16Part1(data)
	if actual != expected {
		t.Errorf("Test Day 16 Part 1 Test 3")
	}

	data = Utils.ReadFileByCharInt("\\inputs_test3.txt")
	expected = "52432133"
	actual = Day16Part1(data)
	if actual != expected {
		t.Errorf("Test Day 16 Part 1 Test 4")
	}

	data = Utils.ReadFileByCharInt("\\inputs_test4.txt")
	expected = "23845678"
	actual = Day16Part1(data)
	if actual != expected {
		t.Errorf("Test Day 16 Part 1 Test 5")
	}
}

func TestDay16Part2(t *testing.T) {
	data := Utils.ReadFileByCharInt("\\inputs.txt")
	expected := "48776785"
	actual := Day16Part2(data)
	if actual != expected {
		t.Errorf("Test Day 16 Part 2 Test 1")
	}

	data = Utils.ReadFileByCharInt("\\inputs_test5.txt")
	expected = "84462026"
	actual = Day16Part2(data)
	if actual != expected {
		t.Errorf("Test Day 16 Part 2 Test 2")
	}

	data = Utils.ReadFileByCharInt("\\inputs_test6.txt")
	expected = "78725270"
	actual = Day16Part2(data)
	if actual != expected {
		t.Errorf("Test Day 16 Part 2 Test 3")
	}

	data = Utils.ReadFileByCharInt("\\inputs_test7.txt")
	expected = "53553731"
	actual = Day16Part2(data)
	if actual != expected {
		t.Errorf("Test Day 16 Part 2 Test 4")
	}
}
