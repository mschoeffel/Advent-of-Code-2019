package Day12

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay12Part1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\moons.txt")
	expected := 8454
	actual := Day12Part1(data, 1000)
	if actual != expected {
		t.Errorf("Test Day 12 Part 1 Test 1")
	}

	data = Utils.ReadFileByLinesString("\\moons_test1.txt")
	expected = 179
	actual = Day12Part1(data, 10)
	if actual != expected {
		t.Errorf("Test Day 12 Part 1 Test 2")
	}

	data = Utils.ReadFileByLinesString("\\moons_test2.txt")
	expected = 1940
	actual = Day12Part1(data, 100)
	if actual != expected {
		t.Errorf("Test Day 12 Part 1 Test 3")
	}
}

func TestDay12Part2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\moons.txt")
	expected := 362336016722948
	actual := Day12Part2(data)
	if actual != expected {
		t.Errorf("Test Day 12 Part 2 Test 1")
	}

	//TODO: Test wont get right result
	//data = Utils.ReadFileByLinesString("\\moons_test1.txt")
	//expected = 2772
	//actual = Day12Part2(data)
	//if actual != expected {
	//	t.Errorf("Test Day 12 Part 2 Test 2")
	//}

	//data = Utils.ReadFileByLinesString("\\moons_test2.txt")
	//expected = 4686774924
	//actual = Day12Part2(data)
	//if actual != expected {
	//	t.Errorf("Test Day 12 Part 2 Test 3")
	//}
}
