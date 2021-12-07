package main

import (
	"fmt"
	"strconv"
)

//判断两条线是否平行

func main() {
	fmt.Println("==================================================================================")
	fmt.Println("Please input the X and Y coordinates for 2 lines to determine if they are parallel")
	fmt.Println("==================================================================================")
	fmt.Println("First Line :")
	fmt.Println("--------------")
	x11, x12, y11, y12 := getCoordinates()
	fmt.Println("Second Line :")
	fmt.Println("--------------")
	x21, x22, y21, y22 := getCoordinates()
	fmt.Println("==================================================================================")
	fmt.Print("Result: ")
	gradient1 := calGradient(x11, x12, y11, y12)
	gradient2 := calGradient(x21, x22, y21, y22)

	if gradient1 == gradient2 {
		fmt.Println("They are parallel")
	} else {
		fmt.Println("They are NOT parallel")
	}

}

func calGradient(x1 float64, x2 float64, y1 float64, y2 float64) (gradient float64) {
	return (y2 - y1) / (x2 - x1)

}

func getCoordinates() (x1 float64, x2 float64, y1 float64, y2 float64) {

	x1 = getCoordinate("start point of X coordinate:")
	y1 = getCoordinate("start point of Y coordinate:")

	x2 = getCoordinate("end point of X coordinate:")
	y2 = getCoordinate("end point of Y coordinate:")

	return

}

func getCoordinate(message string) (coordinate float64) {
	var input string
	var err error

	for {
		fmt.Print(message)
		fmt.Scanln(&input)
		coordinate, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid format. Please input a valid coordinate number.")
		} else {
			break
		}
	}
	return
}
