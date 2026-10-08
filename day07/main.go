package main

import (
	"fmt"
	"log"
	"slices"

	"github.com/kevin-kho/aoc-utilities/common"
)

func SolvePartOne(intArr []int) int {
	// The median will minimize absolute differences in the array
	slices.Sort(intArr)

	var targetVal int
	// Case: odd
	if len(intArr)%2 == 1 {
		mid := len(intArr) / 2
		targetVal = intArr[mid]
	} else {
		// case: even
		r := len(intArr) / 2
		l := r - 1

		targetVal = (intArr[l] + intArr[r]) / 2
	}

	var res int

	for _, val := range intArr {
		res += common.IntAbs(targetVal - val)
	}

	return res

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	intArr, err := common.ParseIntArray(data)
	if err != nil {
		log.Fatal(err)
	}

	res := SolvePartOne(intArr)
	fmt.Println(res)

}
