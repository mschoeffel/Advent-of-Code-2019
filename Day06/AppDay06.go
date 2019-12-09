package Day06

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"strconv"
	"strings"
)

type Planet struct {
	name     string
	children []*Planet
	parent   *Planet
}

func Main() {
	data := Utils.ReadFileByLinesString("\\Day06\\orbits.txt")
	fmt.Println("Result part one: " + strconv.Itoa(Day6Part1(data)))

	data = Utils.ReadFileByLinesString("\\Day06\\orbits.txt")
	fmt.Println("Result part two: " + strconv.Itoa(Day6Part2(data)))
}

func Day6Part1(data []string) int {
	planets := createDataStructure(data)

	steps := 0
	for i := range planets {
		planet := planets[i]
		steps += evaluateParents(planet)
	}
	return steps
}

func Day6Part2(data []string) int {
	planets := createDataStructure(data)
	return evaluateDistance(planets["YOU"], planets["SAN"])
}

func createDataStructure(data []string) map[string]*Planet {
	planets := make(map[string]*Planet)
	for i := range data {
		connection := strings.Split(data[i], ")")
		planetFromName := connection[0]
		planetToName := connection[1]

		planetFrom, ok := planets[planetFromName]
		if !ok {
			planetFrom = &Planet{name: planetFromName, children: nil, parent: nil}
			planets[planetFromName] = planetFrom
		}

		planetTo, ok := planets[planetToName]
		if !ok {
			planetTo = &Planet{name: planetToName, children: nil, parent: nil}
			planets[planetToName] = planetTo
		}

		planetFrom.children = append(planetFrom.children, planetTo)
		planetTo.parent = planetFrom
	}
	return planets
}

func evaluateParents(planet *Planet) int {
	if planet.parent != nil {
		return evaluateParents(planet.parent) + 1
	}
	return 0
}

func evaluateDistance(planetFrom *Planet, planetTo *Planet) int {
	visited := pathToRoot(planetFrom)

	currentPlanet := planetTo
	found := false
	samePlanet := Planet{}
	steps := -1
	for !found {
		if contains(visited, currentPlanet) {
			found = true
			samePlanet = *currentPlanet
			break
		}
		if currentPlanet.parent != nil {
			currentPlanet = currentPlanet.parent
		}
		steps += 1
	}

	foundFrom := false
	currentFromPlanet := planetFrom
	for !foundFrom {
		if currentFromPlanet.name == samePlanet.name {
			foundFrom = true
			break
		}
		if currentFromPlanet.parent != nil {
			currentFromPlanet = currentFromPlanet.parent
		}
		steps += 1
	}
	return steps - 1
}

func pathToRoot(planet *Planet) []string {
	rootFound := false
	var visited []string
	currentPlanet := planet
	for !rootFound {
		if currentPlanet.parent == nil {
			visited = append(visited, currentPlanet.name)
			rootFound = true
			break
		} else {
			visited = append(visited, currentPlanet.name)
			currentPlanet = currentPlanet.parent
		}
	}
	return visited
}

func contains(s []string, e *Planet) bool {
	for _, a := range s {
		if a == e.name {
			return true
		}
	}
	return false
}
