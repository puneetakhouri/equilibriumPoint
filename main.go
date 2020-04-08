package main

import (
	"fmt"
	"strconv"
)

func main() {

	arr := []int{1, 3, 5, 2, 2}
	equilibriumPoint := FindEquilibrium(arr)

	if equilibriumPoint >= len(arr) {
		fmt.Println(strconv.Itoa(-1))
	} else {
		fmt.Println(strconv.Itoa(equilibriumPoint))
	}

}

//FindEquilibrium does the finding
func FindEquilibrium(arr []int) int {
	equilibriumPoint := 0
	leftSum, rightSum := 0, 0
	for i := 0; i < len(arr); i++ {
		rightSum = rightSum + arr[i]
	}

	for equilibriumPoint < len(arr) {
		if leftSum == (rightSum - arr[equilibriumPoint]) {
			equilibriumPoint = equilibriumPoint + 1
			break
		}
		leftSum = leftSum + arr[equilibriumPoint]
		rightSum = rightSum - arr[equilibriumPoint]
		equilibriumPoint = equilibriumPoint + 1
	}
	return equilibriumPoint
}
