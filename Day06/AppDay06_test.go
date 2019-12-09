package Day06

import (
	"Advent-of-Code-2019/Utils"
	"testing"
)

func TestDay06Part1(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\orbits.txt")
	expected := 241064
	actual := Day6Part1(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 1 Test 1")
	}

	data = []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	expected = 42
	actual = Day6Part1(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 1 Test 2")
	}
}

func TestDay06Part2(t *testing.T) {
	data := Utils.ReadFileByLinesString("\\orbits.txt")
	expected := 418
	actual := Day6Part2(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 2 Test 1")
	}

	data = []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"}
	expected = 4
	actual = Day6Part2(data)
	if actual != expected {
		t.Errorf("Test Day 06 Part 1 Test 2")
	}
}
