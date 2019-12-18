package Day18

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"strconv"
	"unicode"
)

type Position struct {
	x     int
	y     int
	value int
}

func Main() {
	maze := Utils.ReadFileByLinesString("//Day18//maze.txt")
	fmt.Println("Result part one: " + strconv.Itoa(Day18Part1(maze)))
}

func Day18Part1(mazeLines []string) int {
	maze := make([][]string, len(mazeLines))
	for i := range maze {
		maze[i] = make([]string, len(mazeLines[0]))
	}

	for i, v := range mazeLines {
		runes := []rune(v)
		for j, r := range runes {
			maze[i][j] = string(r)
		}
	}

	keys := scanMazeKeys(maze)
	doors := scanMazeDoors(maze)
	position := findPlayerPosition(maze)

	dist := findMinPath(position, keys["j"], maze)

	fmt.Println(keys)
	fmt.Println(doors)
	fmt.Println(position)

	return dist
}

func findMinPath(pos Position, target Position, maze [][]string) int {
	pos.value = 0
	possibleMoves := []Position{pos}
	var visited []Position
	for len(possibleMoves) > 0 {
		currentPos := possibleMoves[0]
		possibleMoves = possibleMoves[1:]

		//Check if current pos is target
		if currentPos.x == target.x && currentPos.y == target.y {
			return currentPos.value
		}

		//Calculate next possible possibleMoves
		//LEFT
		if currentPos.x-1 >= 0 {
			posLeft := Position{x: currentPos.x - 1, y: currentPos.y, value: currentPos.value + 1}
			if maze[posLeft.y][posLeft.x] != "#" && !contains(visited, posLeft) {
				possibleMoves = append(possibleMoves, posLeft)
			}
		}
		//RIGHT
		if currentPos.x+1 < len(maze[0]) {
			posRight := Position{x: currentPos.x + 1, y: currentPos.y, value: currentPos.value + 1}
			if maze[posRight.y][posRight.x] != "#" && !contains(visited, posRight) {
				possibleMoves = append(possibleMoves, posRight)
			}
		}
		//DOWN
		if currentPos.y+1 < len(maze) {
			posDown := Position{x: currentPos.x, y: currentPos.y + 1, value: currentPos.value + 1}
			if maze[posDown.y][posDown.x] != "#" && !contains(visited, posDown) {
				possibleMoves = append(possibleMoves, posDown)
			}
		}
		//UP
		if currentPos.y-1 >= 0 {
			posUp := Position{x: currentPos.x, y: currentPos.y - 1, value: currentPos.value + 1}
			if maze[posUp.y][posUp.x] != "#" && !contains(visited, posUp) {
				possibleMoves = append(possibleMoves, posUp)
			}
		}
		visited = append(visited, currentPos)
	}
	return -1
}

func contains(list []Position, elem Position) bool {
	for _, pos := range list {
		if pos.x == elem.x && pos.y == elem.y {
			return true
		}
	}
	return false
}

func scanMazeKeys(maze [][]string) map[string]Position {
	keys := map[string]Position{}
	for i := range maze {
		for j, v := range maze[i] {
			r := []rune(v)
			if unicode.IsLower(r[0]) {
				position := Position{x: j, y: i}
				keys[v] = position
			}
		}
	}
	return keys
}

func scanMazeDoors(maze [][]string) map[string]Position {
	doors := map[string]Position{}
	for i := range maze {
		for j, v := range maze[i] {
			r := []rune(v)
			if unicode.IsUpper(r[0]) {
				position := Position{x: j, y: i}
				doors[v] = position
			}
		}
	}
	return doors
}

func findPlayerPosition(maze [][]string) Position {
	for i := range maze {
		for j, v := range maze[i] {
			if v == "@" {
				return Position{x: j, y: i}
			}
		}
	}
	return Position{-1, -1, 0}
}
