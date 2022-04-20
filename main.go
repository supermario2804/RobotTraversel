package main

import (
	"bufio"
	"log"
	"os"
	"fmt"

	"io"
	"strconv"
	"strings"
)
//a map which will keep track of the travelled point
var travelledPointMap = map[string]string{}

//This would store the value of plane size
var planeSizeCoordinate coordinate

type direction int

const (
	N direction = iota
	E
	S
	W
)

type coordinate struct {
	x int64
	y int64
}

func main() {

	arr := [][]string{}
	file, err := os.Open("input.txt")
	reader := bufio.NewReader(file)
	handleErr(err)

	for i := 0; i < 3; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		arr = append(arr, arrRowTemp)

	}

	//To collect Current Plane size
	planeCoX, err := strconv.ParseInt(arr[0][0], 10, 64)
	handleErr(err)
	planeCoY, err := strconv.ParseInt(arr[0][1], 10, 64)
	handleErr(err)
	planeSizeCoordinate = coordinate{planeCoX, planeCoY}
	//log.Println(planeSizeCoordinate)

	//To Robots current Position
	PositionCoX, err := strconv.ParseInt(arr[1][0], 10, 64)
	handleErr(err)
	PositionCoY, err := strconv.ParseInt(arr[1][1], 10, 64)
	handleErr(err)
	robotInitPos := coordinate{PositionCoX, PositionCoY}
	robotInitDir := (strToDir(arr[1][1]))
	//fmt.Println(robotInitDir, robotInitPos)

	//Gether the moving steps
	stepsToMove := strings.Split(arr[2][0], "")
//	fmt.Println(stepsToMove)
	endpoint, direction :=  returnEndPoint(robotInitPos, robotInitDir, stepsToMove)
	fmt.Println(endpoint, direction)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func returnEndPoint(robotInitPos coordinate, robotInitDir direction, stepsToMove []string) (coordinate, direction) {
	updatedCoordinate, curCoordinate := robotInitPos, robotInitPos
	curDir := robotInitDir
	//travelledPointMap := make(map[string]string)

	for _, move := range stepsToMove {


		switch move {
		case "L":
			curDir = curDir - 1
			curDir = adjustDir(int(curDir))

		case "R":
			curDir = curDir + 1
			curDir = adjustDir(int(curDir))

		case "M":
			if curDir == 0 {
				updatedCoordinate.y = updatedCoordinate.y + 1
			} else if curDir == 1 {
				updatedCoordinate.x = updatedCoordinate.x + 1
			} else if curDir == 2 {
				updatedCoordinate.y = updatedCoordinate.y - 1
			} else if curDir == 3 {
				updatedCoordinate.x = updatedCoordinate.x - 1
			}

		}
	if isValidCoordinate(updatedCoordinate) {
		curCoordinate = updatedCoordinate
		travelledPointMap[fmt.Sprintf("%v",updatedCoordinate)] = "yes"
	} else {
		return curCoordinate, curDir
	}
	}

	/*switch letter {
	case L:
		currDirection = Direction(letter) + 1
	//case R:
	//case M:

	}*/

	 
	
	
	return curCoordinate, curDir
}

func handleErr(err error, args ...interface{}) {
	if err != nil {
		for _, arg := range args {
			log.Printf("%v", arg)
		}
		log.Panic(err)
	}
}

func strToDir(dirStr string) direction {
	switch dirStr {
	case "N":
		return 0
	case "E":
		return 1
	case "S":
		return 2
	case "W":
		return 3
	default:
		return 0
	}
}

func adjustDir(dirInt int) direction {
	if dirInt >= 0 && dirInt <= 3 {
		return direction(dirInt)
	}

	if dirInt == -1 {
		return 3
	}

	if dirInt == 4 {
		return 0
	}
	return direction(dirInt)
}

func isValidCoordinate(co coordinate) bool {
	if co.x > planeSizeCoordinate.x || co.y > planeSizeCoordinate.y || co.x < 0 || co.y < 0 {
		return false
	}

	_, ok := travelledPointMap[fmt.Sprintf("%v", co)]
	if ok {
		return false
	}
	return true
}

func (dir direction) String() string {
	switch dir {
	case 0:
		return "N"
	case 1:
		return "E"
	case 2:
		return "S"
	case 3:
		return "w"
	default:
		return "N"

	}
}