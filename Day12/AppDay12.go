package Day12

import (
	"Advent-of-Code-2019/Utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Moon struct {
	x  int
	y  int
	z  int
	vx int
	vy int
	vz int
}

func Main() {
	moons := Utils.ReadFileByLinesString("//Day12//moons.txt")
	fmt.Println("Result part one: " + strconv.Itoa(Day12Part1(moons, 1000)))

	moons = Utils.ReadFileByLinesString("//Day12//moons.txt")
	fmt.Println("Result part two: " + strconv.Itoa(Day12Part2(moons)))
}

func Day12Part1(moons []string, steps int) int {
	moonList := parseInput(moons)

	for i := 0; i < steps; i++ {
		moonList = calcVelocityMoons(moonList)
		moonList = calcGravityMoons(moonList)
		moonList = resetVelocityMoons(moonList)
	}

	return calcEnergyOfMoons(moonList)
}

func Day12Part2(moons []string) int {
	moonList := parseInput(moons)
	running := true
	steps := 1
	cycle := [3]int{}
	found := [3]bool{}
	initPos := [4][3]int{}
	initPos[0][0] = moonList[0].x
	initPos[0][1] = moonList[0].y
	initPos[0][2] = moonList[0].z

	initPos[1][0] = moonList[1].x
	initPos[1][1] = moonList[1].y
	initPos[1][2] = moonList[1].z

	initPos[2][0] = moonList[2].x
	initPos[2][1] = moonList[2].y
	initPos[2][2] = moonList[2].z

	initPos[3][0] = moonList[3].x
	initPos[3][1] = moonList[3].y
	initPos[3][2] = moonList[3].z

	for running {
		moonList = calcVelocityMoons(moonList)
		moonList = calcGravityMoons(moonList)
		moonList = resetVelocityMoons(moonList)

		if !found[0] &&
			moonList[0].x == initPos[0][0] && moonList[1].x == initPos[1][0] && moonList[2].x == initPos[2][0] && moonList[3].x == initPos[3][0] &&
			moonList[0].vx == 0 && moonList[1].vx == 0 && moonList[2].vx == 0 && moonList[3].vx == 0 {
			found[0] = true
			cycle[0] = steps
		}

		if !found[1] &&
			moonList[0].y == initPos[0][1] && moonList[1].y == initPos[1][1] && moonList[2].y == initPos[2][1] && moonList[3].y == initPos[3][1] &&
			moonList[0].vy == 0 && moonList[1].vy == 0 && moonList[2].vy == 0 && moonList[3].vy == 0 {
			found[1] = true
			cycle[1] = steps
		}

		if !found[2] &&
			moonList[0].z == initPos[0][2] && moonList[1].z == initPos[1][2] && moonList[2].z == initPos[2][2] && moonList[3].z == initPos[3][2] &&
			moonList[0].vz == 0 && moonList[1].vz == 0 && moonList[2].vz == 0 && moonList[3].vz == 0 {
			found[2] = true
			cycle[2] = steps
		}

		if found[0] && found[1] && found[2] {
			running = false
		}
		steps++
	}
	mul := gcf(gcf(cycle[0], cycle[1]), cycle[2])
	return cycle[0] * (cycle[1] / mul) * (cycle[2] / mul)
}

func parseInput(moons []string) []Moon {
	var moonList []Moon
	for i := range moons {
		moonClean := strings.ReplaceAll(moons[i], ">", "")
		moonTemp := strings.Split(moonClean, ", ")
		xTemp, err := strconv.ParseInt(strings.Split(moonTemp[0], "=")[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		yTemp, err := strconv.ParseInt(strings.Split(moonTemp[1], "=")[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		zTemp, err := strconv.ParseInt(strings.Split(moonTemp[2], "=")[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		moonList = append(moonList, Moon{x: int(xTemp), y: int(yTemp), z: int(zTemp), vx: 0, vy: 0, vz: 0})
	}
	return moonList
}

func printMoons(moons []Moon) {
	for i := range moons {
		fmt.Print(" Moon: " + strconv.Itoa(i) + " : " + strconv.Itoa(moons[i].x) + "," + strconv.Itoa(moons[i].y) + "," + strconv.Itoa(moons[i].z))
	}
	fmt.Println()
}

func calcGravity(moon Moon) Moon {
	moon.x += moon.vx
	moon.y += moon.vy
	moon.z += moon.vz
	return moon
}

func calcGravityMoons(moons []Moon) []Moon {
	for i := range moons {
		moons[i] = calcGravity(moons[i])
	}
	return moons
}

func resetVelocity(moon Moon) Moon {
	moon.vx = 0
	moon.vy = 0
	moon.vz = 0
	return moon
}

func resetVelocityMoons(moons []Moon) []Moon {
	for i := range moons {
		resetVelocity(moons[i])
	}
	return moons
}

func calcVelocity(moon1 Moon, moon2 Moon) (Moon, Moon) {
	if moon1.x > moon2.x {
		moon1.vx -= 1
		moon2.vx += 1
	} else if moon1.x < moon2.x {
		moon1.vx += 1
		moon2.vx -= 1
	}
	if moon1.y > moon2.y {
		moon1.vy -= 1
		moon2.vy += 1
	} else if moon1.y < moon2.y {
		moon1.vy += 1
		moon2.vy -= 1
	}
	if moon1.z > moon2.z {
		moon1.vz -= 1
		moon2.vz += 1
	} else if moon1.z < moon2.z {
		moon1.vz += 1
		moon2.vz -= 1
	}
	return moon1, moon2
}

func calcVelocityMoons(moons []Moon) []Moon {
	moons[0], moons[1] = calcVelocity(moons[0], moons[1])
	moons[0], moons[2] = calcVelocity(moons[0], moons[2])
	moons[0], moons[3] = calcVelocity(moons[0], moons[3])
	moons[1], moons[2] = calcVelocity(moons[1], moons[2])
	moons[1], moons[3] = calcVelocity(moons[1], moons[3])
	moons[2], moons[3] = calcVelocity(moons[2], moons[3])
	return moons
}

func calcEnergyOfMoons(moons []Moon) int {
	sum := 0
	for i := range moons {
		sum += calcPotentialEnergy(moons[i]) * calcKineticEnergy(moons[i])
	}
	return sum
}

func calcPotentialEnergy(moon Moon) int {
	xTemp := moon.x
	if xTemp < 0 {
		xTemp = xTemp * -1
	}
	yTemp := moon.y
	if yTemp < 0 {
		yTemp = yTemp * -1
	}
	zTemp := moon.z
	if zTemp < 0 {
		zTemp = zTemp * -1
	}
	return xTemp + yTemp + zTemp
}

func calcKineticEnergy(moon Moon) int {
	vxTemp := moon.vx
	if vxTemp < 0 {
		vxTemp = vxTemp * -1
	}
	vyTemp := moon.vy
	if vyTemp < 0 {
		vyTemp = vyTemp * -1
	}
	vzTemp := moon.vz
	if vzTemp < 0 {
		vzTemp = vzTemp * -1
	}
	return vxTemp + vyTemp + vzTemp
}

func gcf(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcf(b, a%b)
}
