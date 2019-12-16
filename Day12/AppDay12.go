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
}

func Day12Part1(moons []string, steps int) int {
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

	for i := 0; i < steps; i++ {
		moonList = calcVelocityMoons(moonList)
		moonList = calcGravityMoons(moonList)
		//printMoons(moonList)
		moonList = resetVelocityMoons(moonList)
	}

	return calcEnergyOfMoons(moonList)
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
